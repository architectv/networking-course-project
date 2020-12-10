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
  import Navbar from './components/Navbar.svelte';
  import Main from './components/Main.svelte';
  import Docs from './components/Docs.svelte';
  import Status from './components/Status.svelte';
  import Admin from './components/Admin.svelte';
  let dark_theme = true;
  $: theme = dark_theme ? "dark-theme" : "light-theme";
  $: document.body.className = theme;
  let section;
  let reload;
</script>
