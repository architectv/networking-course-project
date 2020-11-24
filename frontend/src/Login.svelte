
<div>
  {#if $user.authorized}
    <Card style="margin: auto; width: max-content;">
      <Content>
        Logged in as {$user.username}
      </Content>
    </Card>
  {:else}
    <Card style="margin: auto; width: max-content;">
      <Content>
        <FormField>
        <Checkbox bind:checked={have_account} />
        <Label>Have account</Label>
        </FormField>
        <br>
        <Textfield bind:value={data.username} label="Username" input$autocomplete="password"/>
        <br />
        <Textfield type="password" bind:value={data.password} label="Password"  />
        {#if !have_account}
          <br />
          <Textfield bind:value={data.dogname} label="Dog name"  />

        {/if}
      </Content>
      <Actions>
        {#if have_account}
          <Button on:click={loginUser}>
            <Label>
              Login
            </Label>
            <Icon class="material-icons">lock</Icon>
          </Button>
        {:else}
          <Button on:click={registerUser}>
            <Label>
              Register
            </Label>
          </Button>
        {/if}
      </Actions>
    </Card>
  {/if}
</div>


<script>
  import Checkbox from '@smui/checkbox';
  import FormField from '@smui/form-field';
  import Textfield from '@smui/textfield';
  import Button, {Icon, Label} from '@smui/button';
  import Card, {Content, Actions} from '@smui/card';
  import {getUser} from './auth';

  let user = getUser();
  $: userJson = JSON.stringify($user);
  let have_account = true;
  let data = {
      username: "",
      password: "",
      dogname: "",
  };
  function loginUser() {
      user.login(data);
  }
  function registerUser() {
      user.register(data);
  }
</script>
