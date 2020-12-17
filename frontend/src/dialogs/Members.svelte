<Dialog bind:this={addMemberDialog}>
  <Title>Add member</Title>
  <Content style="display: flex; flex-direction: column;">
    <TextField bind:value={nickname} label="Nickname" type="text"/>
    <Select enhanced bind:value={role} label="Role">
    {#each roles as r}
      <Option value={r} selected={role === r}>{r}</Option>
    {/each}
    </Select>
  </Content>
  <Actions>
    <Button on:click={() => {members.addMember(nickname, role)}}
    >Add member</Button>
    <Button>Close</Button>
  </Actions>
</Dialog>


<Dialog bind:this={membersDialog}>
  <Title>Members</Title>
  <Content>
  {#each $members as group}
  <div>
  <div class="mdc-typography--headline4">
    {group.name}
  </div>
    {#each group.items as member}
    <div style="display: flex; flex-direction: row; align-items: center;">
      {#if !member.avatar}
        <img src="unknown.png" alt="Avatar" style="border-radius: 50%; height: 48px;">
      {:else}
        <img src="{member.avatar}" alt="Avatar" style="border-radius: 50%; height: 48px;">
      {/if}
      <div style="margin-right: 1em; margin-left: 5px;">
        {member.nickname}
      </div>
    </div>
    {/each}
  </div>
  {/each}
  </Content>
  <Actions>
    <Button on:click={openAddMember}>
      <Label>Add member</Label>
    </Button>
    <Button on:click={() => {}}>
      <Label>Close</Label>
    </Button>
  </Actions>
</Dialog>

<script>
  import Select, {Option} from '@smui/select';
  import TextField from '@smui/textfield';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label} from '@smui/button';
  import {getMembers} from '../api/members';
  export let path;
  let addMemberDialog;
  let nickname = "";
  let roles = ["reader", "writer", "admin"];
  let role = roles[0];
  let members = getMembers(path);
  
  function openAddMember() {
    addMemberDialog.open();
  }
  
  let membersDialog;
  export function open(...args) {
    membersDialog.open(...args);
  }
</script>