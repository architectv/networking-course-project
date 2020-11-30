
<div>
  {#if $user.authorized}
    <div style="margin: auto; width: 80%; height: 80%;">
      <Projects/>
    </div>
  {:else}
    <Login/>
  {/if}
</div>


<script>
  import Card, {Content, Actions} from '@smui/card';
  import {getUser} from './auth';
  import Projects from './Projects.svelte';
  import Login from './Login.svelte';
  let user = getUser();
  $: userJson = JSON.stringify($user);
</script>

