<Dialog bind:this={newDialog}>
  <Title>Dialog Title</Title>
  <Content>
    Do you want create a project?
        <br />
        {#each fields as f, i}
            <Textfield bind:invalid={f.invalid}
                       bind:value={f.value} 
                       on:input={(e) => {validateField(projects, f, e)}}
                       useNativeValidation={false}
                       label={f.name} 
                       textarea={f.long}
                       type={f.type} />
            {#if f.invalid}
            <HelperText validationMsg>{f.error}</HelperText>
            {/if}
            <br />
        {/each}
  </Content>
  <Actions>
    <Button on:click={createProject}>
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
{#if !$projects.current}
<div class="mdc-typography--headline4">
  Projects
  <IconButton on:click={() => newDialog.open()}>
    <Icon class="material-icons">add_circle_outline</Icon>
  </IconButton>
</div>
{/if}


{#if $projects.current}
  <Boards project={project_curr}/>
{:else if $projects.list} 
<List twoLine>
  {#each $projects.list as project, i}
    <Item on:click={() => {project_curr = project; projects.setCurrent(project.id);}}>
      <Text>
        <PrimaryText>{project.title}</PrimaryText>
        {#if project.description}
          <SecondaryText style="width: 50%;">{project.description}</SecondaryText>
        {/if}
      </Text>
      <Meta>
        Changed
       {getDate(project.datetimes.updated)}
      </Meta>
    </Item>
  {/each}
</List>
{:else}
  <p>
  You don't have projects now
  </p>
{/if}
</div>



<script>
  import IconButton from '@smui/icon-button';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label, Icon} from '@smui/button';
  import DataTable, {Head, Body, Row, Cell} from '@smui/data-table';
  import {projects} from '../api/projects.js';
  import HelperText from '@smui/textfield/helper-text/index';
  import Textfield from '@smui/textfield';
  import {validateField, getValidData, getDate} from '../utils';
  import Boards from './Boards.svelte';
  import {onDestroy} from 'svelte';
  import List, {Group, Item, Meta, Separator, Subheader, Text, PrimaryText, SecondaryText} from '@smui/list';
  let project_curr;

  onDestroy(() => {
    projects.release();
  });

  let newDialog;
  let errorDialog;

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

  function createProject() {
    let data = getValidData(fields, projects);
    if (!data) return;
    projects.create(data, (x) => {errorDialog.open();});
  }
</script>
