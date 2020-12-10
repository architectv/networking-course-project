<Dialog bind:this={newDialog}>
  <Title>Dialog Title</Title>
  <Content>
    Do you want create a task?
        <br />
        {#each fields as f, i}
            <Textfield bind:invalid={f.invalid}
                       bind:value={f.value} 
                       on:input={(e) => {validateField(tasks, f, e)}}
                       useNativeValidation={false}
                       label={f.name} 
                       textarea={f.long}
                       type={f.type} />
            {#if f.invalid}
            <HelperText validationMsg>{f.error}</HelperText>
            {/if}
            <br />
        {/each}
        {@html marked(description)}
        <br />
  </Content>
  <Actions>
    <Button on:click={createTask}>
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
    New task wasn't created
  </Content>
</Dialog>

<TaskDialog bind:this={taskDialog} />

{#if tlist}
<List twoLine>
<section use:dndzone={{items: tlist, flipDurationMs}} on:consider={consider} on:finalize={finalize} 
         style="min-height: 2em;">
  {#each tlist as task(task.id)}
  <div animate:flip={{duration: flipDurationMs}}>
  <Item on:click={() => {taskDialog.open(task, () => {deleteTask(task.id)})}}>
    <Text>
      <PrimaryText>{task.title}</PrimaryText>
     <!--<SecondaryText>{JSON.stringify(task)}</SecondaryText>--> 
    </Text>
  </Item>
  </div>
  {/each}
</section>
</List>
{/if}

<Button on:click={() => {newDialog.open()}}>
  <Icon class="material-icons">add_circle_outline</Icon>
  <Label>Add task</Label>
</Button>

<script>
  import { flip } from 'svelte/animate';
  import marked from 'marked';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData} from '../utils';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Icon, Label} from '@smui/button';
  import List, { Item, Text, PrimaryText, SecondaryText, Meta } from '@smui/list';
  import { onDestroy, onMount } from "svelte";
  import { getTasks } from '../api/tasks';
  import {dndzone} from 'svelte-dnd-action';
  import TaskDialog from '../dialogs/TaskDialog.svelte';
  export let listId;
  const flipDurationMs = 300;
  let tasks = getTasks(listId);
  let newDialog;
  let errorDialog;
  let taskDialog;
  export let updateList;
  updateList[listId] = tasks.refresh;
  
  onMount(() => {
    console.log("Mount task", listId);
  })
  
  function deleteTask(id) {
    tasks.deleteTask(id);
  }
  
  function prepareTasks(list) {
    if (!list) return [];
    return list.map((value) => {
      value.id = value._id;
      return value;
    }).sort((a, b) => {
      return a.position < b.position;
    });
  }

  $: tlist = prepareTasks($tasks.list);
  
	function consider(e) {
    tlist = [...e.detail.items];
	}
  
  async function finalize(e) {
    console.log(e);
    for (let index = 0; index < e.detail.items.length; index++) {
      let value = e.detail.items[index];
      let changedId = e.detail.info.id;
      if (value.id == changedId && (value.position != index || listId != value.listId)) {
        let prevListId = value.listId;
        value.position = index;
        value.listId = listId;
        await tasks.updateTask(value.id, prevListId, {
          listId: value.listId, 
          position: value.position});
        if (listId != prevListId && updateList[prevListId]) {
          updateList[prevListId]();
        }
      }
    }
    await tasks.refresh();
  }
  
  export function refresh() {
    tasks.refresh();
  }
  
  onDestroy(() => {
    tasks.release();
  })

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
    
  $: description = fields[1].value;
  
  function createTask() {
    let data = getValidData(fields, tasks);
    if (!data) return;
    tasks.create(data, (x) => {errorDialog.open();});
  }
</script>