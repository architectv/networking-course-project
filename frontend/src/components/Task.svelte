
<TaskDialog bind:this={taskDialog} deleteDialog={deleteDialog} bind:tasks/>

<Paper class="flex-item mdc-elevation--z5" 
       style="width: auto; margin-top: 10px; padding: 0;">
  <div style="display: flex; flex-direction: row">
    <div style="display: flex; align-items: center;">
      <IconButton style="padding-right: 0; width: min-content;" on:click={movePrevList}>
      <Icon class="material-icons">arrow_back_ios</Icon>
      </IconButton>
    </div>
    <div style="display: flex; flex-direction: column;">
      <IconButton style="padding-left: 0; width: min-content;" on:click={decPosition}>
      <Icon class="material-icons">expand_less</Icon>
      </IconButton>
      <IconButton style="padding-left: 0; width: min-content;" on:click={incPosition}>
      <Icon class="material-icons">expand_more</Icon>
      </IconButton>
    </div>
    <div style="width: 100%;"
         on:click={() => {taskDialog.open(task)}} 
    >
      <Title>
        {task.title || "Unnamed"}
      </Title>
      <Content class="flex-container" style="flex-wrap: wrap; max-width: 25vw;">
          {#each task.labels as label}
          <div>
            <TaskLabel label={label} />
          </div>
          {/each}
      </Content>
    </div>
    <div style="display: flex; align-items: center;">
      <IconButton style="padding: 0; width: min-content;" on:click={moveNextList}>
      <Icon class="material-icons">arrow_forward_ios</Icon>
      </IconButton>
    </div>
  </div>
</Paper>

<script>
  import Paper, {Title, Content} from '@smui/paper';
  import IconButton, {Icon} from '@smui/icon-button';
  import {dndzone} from 'svelte-dnd-action';
  import { flip } from 'svelte/animate';
  import TaskLabel from './TaskLabel.svelte';
  import TaskDialog from '../dialogs/TaskDialog.svelte';
  export let task;
  export let tasks;
  export let deleteDialog;
  export let moveNext;
  export let movePrev;
  let taskDialog;
  
  function moveNextList() {
    moveNext(task);
  }
  
  function movePrevList() {
    movePrev(task);
  }
  
  function incPosition() {
    tasks.updateTask(task.id, task.listId, {position: task.position + 1});
  }
  
  function decPosition() {
    tasks.updateTask(task.id, task.listId, {position: task.position - 1});
  }
</script>