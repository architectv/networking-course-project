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

<Dialog bind:this={errorDialog}>
  <Title>Error</Title>
  <Content>
    New list wasn't created
  </Content>
</Dialog>

<div style="margin: auto;">
<h1>
  Lists
  <Button on:click={() => newDialog.open()}>
    <Label>
      New
    </Label>
  </Button>
</h1>
<br/>

{#if $lists.list} 
  <List>
  {#each $lists.list as list, i}
    <Item>
      <Text>{list.title}</Text>
    </Item>
  {/each}
  </List>
{:else}
  <p>
  You don't have lists now
  </p>
{/if}
</div>



<script>
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import DataTable, {Head, Body, Row, Cell} from '@smui/data-table';
  import {lists} from './lists.js';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData} from './utils';
  import Drawer, {Subtitle, Scrim} from '@smui/drawer';
  import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
  import {onDestroy} from 'svelte';

  onDestroy(() => {
    lists.release();
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

  function createList() {
    let data = getValidData(fields, lists);
    if (!data) return;
    lists.create(data, (x) => {errorDialog.open();});
  }
</script>
