
<Card style="width: 50%; margin: 1em auto; min-width: 250px;" padded>
	{@html $data}
</Card>

<Card style="width: 50%; margin: 1em auto; min-width: 250px;" padded>
{#if $user.authorized}
  Logged in as {$user.username}
{:else}
  Unauthorized
{/if}
</Card>

<Fab on:click={reload} bind:exited class="floating">
  <Icon class="material-icons">cached</Icon>
</Fab>

<script>
  import { writable } from "svelte/store";
  import { onMount, onDestroy } from 'svelte';
  import Fab from '@smui/fab';
  import {Icon} from '@smui/common';
  import Card, {Content, PrimaryAction, Media, MediaContent, Actions, ActionIcons} from '@smui/card';
  import {user} from './auth';

  const data = writable("");
  const isFetching = writable(false);
  export let reload;
  let exited = $isFetching || !reload;

  onMount(() => {
      reload = reloadStatus;
      reload();
  });

  onDestroy(() => {
      reload = undefined;
  });

  function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  async function reloadStatus() {
      if ($isFetching) {
          return;
      }
      isFetching.set(true);
      let text = await fetch("status").then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.text();
      }).then((x) => {
        return x;
      }).catch((x) => {
        return "Error" + x;
      });
      data.set(text.replace(/(?:\r\n|\r|\n)/g, '<br>'));
      isFetching.set(false);
  }
</script>

