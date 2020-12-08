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

<Dialog bind:this={errorDialog}>
  <Title>Error</Title>
  <Content>
    New project wasn't created
  </Content>
</Dialog>

<div style="margin: auto;">
{#if !$boards.current}
<h1>
  Boards
  <Button on:click={() => newDialog.open()}>
    <Label>
      New
    </Label>
  </Button>
</h1>
<br/>
{/if}

{#if $boards.current}
  <Lists/>
{:else if $boards.list} 
  <List>
  {#each $boards.list as board, i}
    <Item on:click={() => {boards.setCurrent(board.id)}}>
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
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import DataTable, {Head, Body, Row, Cell} from '@smui/data-table';
  import {boards} from './boards.js';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData} from './utils';
  import Drawer, {Subtitle, Scrim} from '@smui/drawer';
  import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
  import Lists from './Lists.svelte';
  import {onDestroy} from 'svelte';

  onDestroy(() => {
    boards.release();
  });


  let newDialog;
  let errorDialog;

  let fields = [
      {
          name: "Title", key: "title", 
          value: "", 
          type: "text", invalid: false,
          error: ""
      }
    ];

  function createBoard() {
    let data = getValidData(fields, boards);
    if (!data) return;
    boards.create(data, (x) => {errorDialog.open();});
  }
</script>
