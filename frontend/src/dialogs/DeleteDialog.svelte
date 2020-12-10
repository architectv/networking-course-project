<Dialog bind:this={deleteDialog}>
{#if item}
  <Title>Delete {item.title}</Title>
  <Content>
  Do you really want to delete {item.title}?
  </Content>
{/if}
  <Actions>
{#if deleteItem}
    <Button on:click={deleteItem}>
      <Label>Delete</Label>
    </Button>
{/if}
    <Button on:click={() => {}}>
      <Label>Close</Label>
    </Button>
  </Actions>
</Dialog>

<script>
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label} from '@smui/button';
  let item;
  let deleteItem;

  let deleteDialog;
  export function open(_item, _deleteItem, ...args) {
    console.log("Delete", _item)
    item = _item;
    deleteItem = _deleteItem;
    deleteDialog.open(...args);
  }
</script>