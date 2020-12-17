
<Dialog bind:this={errorDialog}>
  <Title>Error</Title>
  <Content>
    New task wasn't created
  </Content>
</Dialog>

{#if tlist}
<section class="flex-container"
         style="min-height: 2em; flex-direction: column;">
  {#each tlist as task(task.id)}
  <div animate:flip={{duration: flipDurationMs}}>
       <Task deleteDialog={deleteDialog} bind:task bind:tasks 
             moveNext={moveNext} movePrev={movePrev}/>
  </div>
  {/each}
</section>
{/if}

<Button on:click={openNewTaskDialog}>
  <Icon class="material-icons">add_circle_outline</Icon>
  <Label>Add task</Label>
</Button>


<script>
  import Chip, {Set, Checkmark} from '@smui/chips';
  import Paper from '@smui/paper';
  import { flip } from 'svelte/animate';
  import {validateField, getValidData} from '../utils';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Icon, Label} from '@smui/button';
  import { onDestroy, onMount } from "svelte";
  import { getTasks } from '../api/tasks';
  import { lists } from '../api/lists';
  import {dndzone} from 'svelte-dnd-action';
  import TaskDialog from '../dialogs/TaskDialog.svelte';
  import Task from './Task.svelte';
  export let project;
  export let board;
  export let list;
  export let newDialog;
  export let deleteDialog;
  export let updateList;

  const flipDurationMs = 300;

  let errorDialog;
  let taskDialog;
  let prevListId;
  let nextListId;
  let listId = list.id;
  let tasks = getTasks(project.id, board.id, listId);

  updateList[listId] = tasks.refresh;
  
  $: ((lst) => {
    prevListId = undefined;
    nextListId = undefined;
    if (!lst) {
      return;
    }
    lst.forEach((val) => {
      if (val.position == list.position - 1) {
        prevListId = val.id;
      } else if (val.position == list.position + 1) {
        nextListId = val.id;
      }
    });
  })($lists.list)
  
  function prepareTasks(list) {
    console.log(list);
    if (!list) return [];
    return list.map((value) => {
      value.id = value._id;
      if (!value.labels) {
        value.labels = [];
      }
      return value;
    }).sort((a, b) => {
      return a.position - b.position;
    });
  }

  $: tlist = prepareTasks($tasks.list);

  export function refresh() {
    tasks.refresh();
  }
  
  onDestroy(() => {
    tasks.release();
  })
  
  function moveNext(task) {
    if (!nextListId) {
      return;
    }
    tasks.updateTask(task.id, listId, {listId: nextListId, position: 0});
    lists.refresh();
  }
  
  function movePrev(task) {
    if (!prevListId) {
      return;
    }
    tasks.updateTask(task.id, listId, {listId: prevListId, position: 0});
    lists.refresh();
  }
  
  let fields = [
      {
          name: "Title", key: "title", 
          value: "", 
          type: "text", invalid: false,
          error: ""
      },
      {
          name: "Description", key: "description", 
          value: "", long: true,
          type: "text", invalid: false,
          error: ""
      },
    ];
    
  function openNewTaskDialog() {
    console.log(newDialog);
    if (!newDialog) return;
    newDialog.open(tasks, fields, 1, createTask, "Create task", "Do you want create task?");
  }
  
  function createTask(fields) {
    let data = getValidData(fields, tasks);
    if (!data) return;
    tasks.create(data, (x) => {errorDialog.open();});
  }
</script>