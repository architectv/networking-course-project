<NewDialog bind:this={updateDialog} />

<Dialog bind:this={labelDialog}>
  <Title>Add label</Title>
  <Content class="flex-container" style="flex-wrap: wrap; max-width: 40vw;">
    {#each newLabels as label}
    <span on:dblclick={() => {tasks.addLabel(task.id, label.id)}}>
      <TaskLabel label={label}/>
    </span>
    {/each}
  </Content>
  <Actions>
    <Button on:click={closeDialog}>
      <Label>Close</Label>
    </Button>
  </Actions>
</Dialog>


<Dialog bind:this={taskDialog}>
{#if task}
  <Title>Task {task.title}
{#if deleteDialog}
  <IconButton on:click={openDeleteDialog}>
    <Icon class="material-icons">delete_outline</Icon>
  </IconButton>
{/if}
{#if updateDialog}
  <IconButton on:click={openUpdateDialog}>
    <Icon class="material-icons">create</Icon>
  </IconButton>
{/if}
  </Title>
  <Content>
  Labels:
  <div class="flex-container" style="flex-wrap: wrap;">
  {#if task.labels}
    {#each task.labels as label}
      <TaskLabel label={label}/>
    {/each}
  {/if}
  </div>
  <IconButton on:click={openLabelsDialog}>
    <Icon class="material-icons">add_outline</Icon>
  </IconButton>

  <br/>
  {#if task.description}
  {@html marked(task.description)}
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
  import NewDialog from '../dialogs/NewDialog.svelte';
  import Button, {Label, Icon} from '@smui/button';
  import IconButton from '@smui/icon-button';
  import TaskLabel from '../components/TaskLabel.svelte';
  import {labels} from '../api/board_labels';
  import { getValidData } from '../utils';
  export let tasks;
  let task;
  
  $: newLabels = ((llist) => {
    if (!llist || !task || !task.labels) {
      return [];
    }
    console.log(llist, task.labels);
    return llist.filter((value) => {
      for (const label of task.labels) {
        if (label.id == value.id) {
          return false;
        }
      }
      return true;
    });
  })($labels.list);
  
  $: console.log(newLabels);

  export let deleteDialog;
  let updateDialog;
  let taskDialog;
  let labelDialog;
  export function open(_item, ...args) {
    task = _item;
    taskDialog.open(...args);
  }
  
  function openLabelsDialog() {
    taskDialog.close();
    labelDialog.open();
  }
  
  function closeDialog() {
  }
  
  function _deleteTask(id) {
    tasks.deleteTask(id);
  }
  
  function openDeleteDialog() {
    if (!deleteDialog) {
      return;
    }
    taskDialog.close();
    deleteDialog.open(task, _deleteTask);
  }

  $: fields = [
      {
          name: "Title", key: "title", 
          value: (task && task.title) || "", 
          type: "text", invalid: false,
          error: ""
      },
      {
          name: "Description", key: "description", 
          value: (task && task.description) || "", long: true,
          type: "text", invalid: false,
          error: ""
      },
    ];
    
  function _updateItem(fields) {
    let data = getValidData(fields, tasks);
    if (!data) return;
    tasks.updateTask(task.id, task.listId, data);
  }
  
  function openUpdateDialog() {
    if (!updateDialog) {
      return;
    }
    taskDialog.close();
    updateDialog.open(tasks, fields, undefined, _updateItem, "Change task", "", true);
  }
</script>