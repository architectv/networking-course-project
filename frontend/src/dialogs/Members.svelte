<Dialog bind:this={membersDialog}>
  <Title>Members</Title>
  <Content>
  {#await getMembers()}
    Loading...
  {:then value}
  {#each value as group}
    {group.name}
    <br/>
    {#each group.items as member}
      {member.nickname}
      <br/>
    {/each}
  {/each}
  {:catch error}
    {error}
  {/await}
  </Content>
  <Actions>
    <Button on:click={() => {}}>
      <Label>Close</Label>
    </Button>
  </Actions>
</Dialog>

<script>
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Label} from '@smui/button';
  export let path;
  async function getMembers() {
    if (!localStorage.token) {
      return [];
    }
    let members = fetch(path, {
      headers: {
        Authorization: `Bearer ${localStorage.token}`
      }
    }).then((resp) => {
      if (resp.ok) {
        return resp.json();
      }
      throw Error("Network error");
    }).then((data) => {
      let m = data.data.members;
      let owner = [];
      let admins = [];
      let writers = [];
      let readers = [];
      m.forEach((value) => {
        if (value.isOwner) owner.push(value);
        else if (value.permissions.admin) admins.push(value);
        else if (value.permissions.write) writers.push(value);
        else readers.push(value);
      });
      let res = [];
      res.push({name: 'Бог', items: owner});
      if (admins.length) res.push({name: 'Админы', items: admins});
      if (writers.length) res.push({name: 'Писатели', items: writers});
      if (readers.length) res.push({name: 'Читатели', items: readers});
      console.log(res);
      return res;
    });
    return members;
  }


  let membersDialog;
  export function open(...args) {
    membersDialog.open(...args);
  }
</script>