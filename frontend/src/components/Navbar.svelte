<TopAppBar dense variant="fixed" style="background: var(--mdc-theme-surface);">
  <Row>
    <Section>
      <IconButton class="material-icons"
                  on:click={() => menu.setOpen(true)}>menu</IconButton>
      <Menu anchorCorner="BOTTOM_LEFT" bind:this={menu}>
        <List style="color: var(--mdc-theme-text-primary-on-background);">
          <Item on:SMUI:action={() => section = 'Main'}><Text>Main</Text></Item>
            <Item on:SMUI:action={() => section = 'Admin'}><Text>Admin</Text></Item>
              <Item on:SMUI:action={() => section = 'Docs'}><Text>Docs</Text></Item>
                <Item on:SMUI:action={() => section = 'Status'}><Text>Status</Text></Item>
        </List>
      </Menu>
      <Title style="color: var(--mdc-theme-text-primary-on-background);">Yak</Title>
    </Section>
    <Section align="end" toolbar>
      {#if $user.data}
        {#if !$user.data.avatar}
          <img src="unknown.png" alt="Avatar" style="border-radius: 50%;">
        {:else}
          <img src="{$user.data.avatar}" alt="Avatar" style="border-radius: 50%;">
        {/if}
        <div style="margin-right: 1em; margin-left: 5px;">
          {$user.data.nickname}
        </div>
      {/if}
      {#if reload}
        <IconButton on:click={reload}>
          <Icon class="material-icons">cached</Icon>
        </IconButton>
      {/if}
      <IconButton toggle bind:pressed={dark_theme}>
        <Icon class="material-icons" on>bedtime</Icon>
        <Icon class="material-icons">brightness_7</Icon>
      </IconButton>
      {#if $user.authorized}
        <IconButton on:click={user.logout}>
          <Icon class="material-icons">directions_run</Icon>
        </IconButton>
      {/if}
    </Section>
  </Row>
</TopAppBar>



<script>
  import TopAppBar, {Row, Section, Title, FixedAdjust} from '@smui/top-app-bar';
  import {onMount, onDestroy} from 'svelte';
  import IconButton from '@smui/icon-button';
  import {Icon} from '@smui/common';
  import Menu from '@smui/menu';
  import List, {Item, Separator, Text, PrimaryText, SecondaryText, Graphic} from '@smui/list';
  import {user} from '../api/auth';
  export let dark_theme;
  export let section="Main";
  export let reload;
  let menu;
</script>

<style>
</style>
