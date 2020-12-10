<Dialog bind:this={newDialog}>
  <Title>Create list</Title>
  <Content>
    Do you want create a list?
        <br />
        {#each fields as f, i}
            <Textfield bind:invalid={f.invalid}
                       bind:value={f.value} 
                       on:input={(e) => {validateField(lists, f, e)}}
                       useNativeValidation={false}
                       label={f.name} 
                       type={f.type} />
            {#if f.invalid}
            <HelperText validationMsg>{f.error}</HelperText>
            {/if}
            <br />
        {/each}
  </Content>
  <Actions>
    <Button on:click={createList}>
      <Label>Create</Label>
    </Button>
    <Button on:click={() => {}}>
      <Label>Cancel</Label>
    </Button>
  </Actions>
</Dialog>

<Members bind:this={membersDialog} path={`api/v1/projects/${project.id}/members`} />

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
</div>



{#if tlists} 
<div class="flex-container">
<section use:dndzone={{items: tlists, type: "columns", flipDurationMs}} on:consider={consider} on:finalize={finalize} 
         style="min-height: 5px;" class="flex-container">
  {#each tlists as list(list.id)}
    <div class="flex-item" animate:flip={{duration: flipDurationMs}}>
      <Paper color="primary" class="list-paper flex-item">
        <Title style="width: max-content;">{list.title}
          <IconButton on:click={() => {deleteDialog.open(list, () => {deleteList(list.id)})}}>
            <Icon class="material-icons">delete_outline</Icon>
          </IconButton>
        </Title>
        {#if !isConsider}
        <Content>
          <Tasks bind:updateList listId={list.id} updateAnotherList={updateList}/>
        </Content>
        {/if}
      </Paper>
    </div>
  {/each}
</section>
<Button on:click={newDialog.open} style="min-width: max-content; margin-left: 5px;">
  <Icon class="material-icons">add_circle_outline</Icon>
  <Label>Add list</Label>
</Button>
</div>
{:else}
  <p>
  You don't have lists now
  </p>
{/if}
</div>


<script>
  import DeleteDialog from '../dialogs/DeleteDialog.svelte';
  import { flip } from 'svelte/animate';
  import {dndzone} from 'svelte-dnd-action';
  import Tasks from './Tasks.svelte';
  import Paper from '@smui/paper';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import IconButton from '@smui/icon-button';
  import {lists} from '../api/lists.js';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData} from '../utils';
  import {onDestroy} from 'svelte';
  import {boards} from '../api/boards';
  import Members from '../dialogs/Members.svelte';

  export let board;
  export let project;
  let isConsider = false;

  const flipDurationMs = 300;
  
  function prepareLists(list) {
    if (!list) return [];
    console.log(list);
    return list.map((value) => {
      return value;
    }).sort((a, b) => {
      return a.position < b.position;
    });
  }
  
  $: tlists = prepareLists($lists.list);
  
	function consider(e) {
    isConsider = true;
		tlists = e.detail.items;
	}
  
  function deleteList(id) {
    lists.deleteList(id);
  }
  
  async function finalize(e) {
    for (let index = 0; index < e.detail.items.length; index++) {
      let value = e.detail.items[index];
      let changedId = e.detail.info.id;
      if (value.id == changedId && (value.position != index)) {
        value.position = index;
        await lists.updateList(value.id, {
          position: value.position
        });
      }
    }
    isConsider = false;
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

  function createList() {
    let data = getValidData(fields, lists);
    if (!data) return;
    lists.create(data, (x) => {errorDialog.open();});
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