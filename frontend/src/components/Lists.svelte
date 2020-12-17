<NewDialog bind:this={newDialog} />

<Members bind:this={membersDialog} path={`api/v1/projects/${project.id}/boards/${board.id}`} />

<Dialog bind:this={errorDialog}>
  <Title>Error</Title>
  <Content>
    New list wasn't created
  </Content>
</Dialog>

<DeleteDialog bind:this={deleteDialog} />

<div style="margin: auto;">
  <div class="mdc-typography--headline4">
    <IconButton on:click={() => {boards.unsetCurrent()}}>
      <Icon class="material-icons">keyboard_backspace</Icon>
    </IconButton>
    {board.title}
    <IconButton on:click={membersDialog.open}>
      <Icon class="material-icons">people</Icon>
    </IconButton>
    <IconButton on:click={() => {deleteDialog.open(board, deleteBoard)}}>
      <Icon class="material-icons">delete_outline</Icon>
    </IconButton>
    <IconButton on:click={openEditBoard}>
      <Icon class="material-icons">create</Icon>
    </IconButton>
  </div>



  {#if tlists} 
    <div class="flex-container">
      <section style="min-height: 20px;" class="flex-container">
        {#each tlists as list(list.id)}
          <div class="flex-item" animate:flip={{duration: flipDurationMs}}>
            <Paper color="primary" class="list-paper flex-item">
              <Title style="width: max-content;">
                <IconButton on:click={() => listMovePrev(list)}>
                  <Icon class="material-icons">arrow_back_ios</Icon>
                </IconButton>
                {list.title}
                <IconButton on:click={() => listMoveNext(list)}>
                  <Icon class="material-icons">arrow_forward_ios</Icon>
                </IconButton>
                <IconButton on:click={openChangeList(list)}>
                  <Icon class="material-icons">create</Icon>
                </IconButton>
                <IconButton on:click={() => {deleteDialog.open(list, () => {deleteList(list.id)})}}>
                  <Icon class="material-icons">delete_outline</Icon>
                </IconButton>
              </Title>
              {#if !isConsider}
              <Content>
                <Tasks bind:updateList list={list} bind:board bind:project
                       newDialog={newDialog} deleteDialog={deleteDialog} />
              </Content>
              {/if}
            </Paper>
          </div>
        {/each}
      </section>
      <Button on:click={openNewListDialog} style="min-width: max-content; margin-left: 5px;">
        <Icon class="material-icons">add_circle_outline</Icon>
        <Label>Add list</Label>
      </Button>
    </div>
  {:else}
    <p>
    You don't have lists now
    </p>
  {/if}

  <div class="mdc-typography--headline5">
    Labels
  </div>

  <div class="flex-container">
    {#each llist as label(label.id)}
      <div class="flex-item" 
           style="margin-top: 10px;"
           animate:flip={{duration: flipDurationMs}}>
           <TaskLabel label={label} 
                  onDelete={() => {deleteDialog.open(label, () => {deleteLabel(label.id)})}} />
      </div>
    {/each}
    <br/>
      <Button on:click={openNewLabelDialog} style="min-width: max-content; margin-left: 5px;">
        <Icon class="material-icons">add_circle_outline</Icon>
        <Label>Add label</Label>
      </Button>
  </div>
</div>


<script>
  import DeleteDialog from '../dialogs/DeleteDialog.svelte';
  import { flip } from 'svelte/animate';
  import {dndzone, SOURCES, TRIGGERS} from 'svelte-dnd-action';
  import Tasks from './Tasks.svelte';
  import Paper from '@smui/paper';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import TaskLabel from './TaskLabel.svelte';
  import IconButton from '@smui/icon-button';
  import {lists} from '../api/lists.js';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData, toColor} from '../utils';
  import {onDestroy} from 'svelte';
  import {boards} from '../api/boards';
  import {labels} from '../api/board_labels';
  import Members from '../dialogs/Members.svelte';
  import NewDialog from '../dialogs/NewDialog.svelte';


  export let board;
  export let project;
  
  let isConsider = false;

  const flipDurationMs = 300;

  let fieldsBoard = [
      {
          name: "Title", key: "title", 
          value: board.title || "", 
          type: "text", invalid: false,
          error: ""
      },
      {
          name: "Read", key: "read", 
          value: board.defaultPermissions.read || false, 
          type: "checkbox", invalid: false,
          error: ""
      },
      {
          name: "Write", key: "write", 
          value: board.defaultPermissions.write || false, 
          type: "checkbox", invalid: false,
          error: ""
      }
    ];
  
  async function updateBoard(fields) {
    let data = getValidData(fields, boards);
    if (!data) return;
    let ret = await boards.updateCurrent(data);
    if (ret) board.title = data.title;
  }
  
  function openEditBoard() {
    newDialog.open(boards, fieldsBoard, undefined, updateBoard, "Change board", "", true);
  }
  
  function openNewListDialog() {
    newDialog.open(lists, fields, undefined, createList, "Create list", "Do you want create list?");
  }
  
  function openNewLabelDialog() {
    newDialog.open(labels, fieldsLabel, undefined, createLabel, "Create label", "Do you want create label?");
  }
  
  function listMoveNext(list) {
    lists.updateList(list.id, {position: list.position + 1});
  }
  
  function listMovePrev(list) {
    lists.updateList(list.id, {position: list.position - 1});
  }
  
  function prepareLists(list) {
    if (!list) return [];
    return list.map((value) => {
      return value;
    }).sort((a, b) => {
      return a.position - b.position;
    });
  }
  
  function prepareLabels(list) {
    if (!list) return [];
    return list
  }
  
  $: tlists = prepareLists($lists.list);
  
  $: llist = prepareLabels($labels.list);

  function deleteList(id) {
    lists.deleteList(id);
  }
  
  function deleteLabel(id) {
    labels.deleteLabel(id);
  }
  
  let updateList = {};
  
  onDestroy(() => {
    lists.release();
  });

  let newDialog;
  let errorDialog;
  let deleteDialog;
  let membersDialog;

  let fields = [
      {
          name: "Title", key: "title", 
          value: "", 
          type: "text", invalid: false,
          error: ""
      }
    ];
    
  function changeListFunction(id) {
    return (fields) => {
      let data = getValidData(fields, lists);
      if (!data) return;
      lists.updateList(id, data);
    }
  }
  
  function openChangeList(list) {
    fields[0].value = list.title;
    newDialog.open(lists, fields, undefined, changeListFunction(list.id), 
                   "Change list", "", true);
  }

  function createList(fields) {
    let data = getValidData(fields, lists);
    if (!data) return;
    lists.create(data, (x) => {errorDialog.open();});
  }

  let fieldsLabel = [
      {
          name: "Name", key: "name", 
          value: "", 
          type: "text", invalid: false,
          error: ""
      },
      {
          type: "color", value: "", invalid: false,
          name: "Color", key: "color", error: ""
      }
    ];

  function createLabel(fields) {
    let data = getValidData(fields, labels);
    if (!data) return;
    labels.create(data, (x) => {errorDialog.open();});
  }

  
  function deleteBoard() {
    boards.deleteCurrent();
  }
</script>

<style>
  .flex-container {
    display: flex;
    flex-direction: row;
  }
  
  .flex-item {
    width: 23%;
  }
</style>