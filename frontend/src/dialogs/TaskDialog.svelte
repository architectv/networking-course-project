<DeleteDialog bind:this={deleteDialog} style="z-index: 2000;" />

<Dialog bind:this={taskDialog}>
{#if item}
  <Title>Task {item.title}
{#if deleteItem}
  <IconButton on:click={openDeleteDialog}>
    <Icon class="material-icons">delete_outline</Icon>
  </IconButton>
{/if}
  </Title>
  <Content>
  {#if item.description}
  {@html marked(item.description)}
  {:else}
  Empty description
  {/if}
  </Content>
{/if}
  <Actions>
    <Button on:click={() => {}}>
      <Label>Close</Label>
    </Button>
  </Actions>
</Dialog>

<script>
  import marked from 'marked';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import DeleteDialog from './DeleteDialog.svelte';
  import IconButton from '@smui/icon-button';
  let item;
  let deleteItem;

  let deleteDialog;
  let taskDialog;
  export function open(_item, _deleteItem, ...args) {
    item = _item;
    deleteItem = _deleteItem;
    taskDialog.open(...args);
  }
  
  function openDeleteDialog() {
    taskDialog.close();
    deleteDialog.open(item, deleteItem);
  }
</script>