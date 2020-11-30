
<div style="margin: auto;">
<h1>
  Projects
</h1>
<br/>

{#if $projects.list} 
<DataTable table$aria-label="People list">
  <Head>
    <Row>
      <Cell>Id</Cell>
      <Cell>Owner Id</Cell>
      <Cell>Title</Cell>
      <Cell>Description</Cell>
    </Row>
  </Head>
  <Body>
    {#each $projects.list as project, i}
      <Row>
        <Cell>{project.id}</Cell>
        <Cell>{project.ownerId}</Cell>
        <Cell>{project.title}</Cell>
        <Cell>{project.description}</Cell>
      </Row>
    {/each}
  </Body>
</DataTable>
{:else}
  <p>
  You don't have projects now
  </p>
{/if}
</div>



<script>
  import DataTable, {Head, Body, Row, Cell} from '@smui/data-table';
  import {getProjects} from './projects.js';
  let projects = getProjects();
</script>
