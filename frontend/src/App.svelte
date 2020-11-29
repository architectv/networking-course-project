<div class="{theme} main">
  <Navbar bind:dark_theme={dark_theme} bind:section={section} bind:reload={reload} />
  <div class="main-content app">
    {#if section == "Docs"}
      <Docs/>
    {:else if section == "Admin"}
      <Admin/>
    {:else if section == "Status"}
      <Status bind:reload={reload} />
    {:else}
      <Main/>
    {/if}
  </div>
</div>

<script>
  import Navbar from './Navbar.svelte';
  import Main from './Main.svelte';
  import Docs from './Docs.svelte';
  import Status from './Status.svelte';
  import Admin from './Admin.svelte';
  let dark_theme = true;
  $: theme = dark_theme ? "dark-theme" : "light-theme";
  let section;
  let reload;
</script>
