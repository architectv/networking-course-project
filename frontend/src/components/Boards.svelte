<Dialog bind:this={newDialog}>
  <Title>Dialog Title</Title>
  <Content>
    Do you want create a board?
        <br />
        {#each fields as f, i}
            <Textfield bind:invalid={f.invalid}
                       bind:value={f.value} 
                       on:input={(e) => {validateField(boards, f, e)}}
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
    <Button on:click={createBoard}>
      <Label>Create</Label>
    </Button>
    <Button on:click={() => {}}>
      <Label>Cancel</Label>
    </Button>
  </Actions>
</Dialog>

<Members bind:this={membersDialog} path={`api/v1/projects/${project.id}`} />

<Dialog bind:this={deleteDialog}>
  <Title>Dialog Title</Title>
  <Content>
    Do you want delete a project?
  </Content>
  <Actions>
    <Button on:click={deleteProject}>
      <Label>Delete</Label>
    </Button>
    <Button on:click={() => {}}>
      <Label>Cancel</Label>
    </Button>
  </Actions>
</Dialog>


<div style="margin: auto;">
{#if !$boards.current}
<div class="mdc-typography--headline4">
  <IconButton on:click={() => {projects.unsetCurrent()}}>
    <Icon class="material-icons">keyboard_backspace</Icon>
  </IconButton>
  {project.title}
  <IconButton on:click={membersDialog.open}>
    <Icon class="material-icons">people</Icon>
  </IconButton>
  <IconButton on:click={deleteDialog.open}>
    <Icon class="material-icons">delete_outline</Icon>
  </IconButton>
</div>

{#if project.description}
<div class="mdc-typography--body1">
    {project.description}
</div>
{/if}


<div class="mdc-typography--headline6">
    Boards
  <IconButton on:click={() => newDialog.open()}>
    <Icon class="material-icons">add_circle_outline</Icon>
  </IconButton>
</div>
<br/>
{/if}

{#if $boards.current}
  <Lists board={board_curr} project={project} />
{:else if $boards.list} 
  <List>
  {#each $boards.list as board, i}
    <Item on:click={() => {board_curr = board; boards.setCurrent(board.id)}}>
      <Text>{board.title}</Text>
    </Item>
  {/each}
  </List>
{:else}
  <p>
  You don't have boards now
  </p>
{/if}
</div>

<script>
  import Members from '../dialogs/Members.svelte';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import IconButton from '@smui/icon-button';
  import DataTable, {Head, Body, Row, Cell} from '@smui/data-table';
  import {boards} from '../api/boards.js';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData} from '../utils';
  import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
  import Lists from './Lists.svelte';
  import {onDestroy} from 'svelte';
  import {projects} from '../api/projects';

  export let project;
  let board_curr;

  onDestroy(() => {
    boards.release();
  });

  let newDialog;
  let membersDialog;
  let errorDialog;
  let deleteDialog;

  let fields = [
      {
          name: "Title", key: "title", 
          value: "", 
          type: "text", invalid: false,
          error: ""
      }
    ];
  
  function deleteProject() {
    projects.deleteCurrent();
  }

  function createBoard() {
    let data = getValidData(fields, boards);
    if (!data) return;
    boards.create(data, (x) => {errorDialog.open();});
  }
</script>
