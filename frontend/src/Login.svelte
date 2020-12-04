
<div>
  {#if $user.authorized}
    <Card style="margin: auto; width: max-content;">
      <Content>
        Logged in as {$user.nickname}
      </Content>
    </Card>
  {:else}
    <Card style="margin: auto; width: max-content;">
      <Content>
        <FormField>
        <Checkbox bind:checked={have_account} />
        <Label>Have account</Label>
        </FormField>
        <br />
        {#each fields as f, i}
          {#if (have_account && !f.register) || !have_account}
            <Textfield bind:invalid={f.invalid}
                       bind:value={f.value} 
                       on:input={(e) => {validateField(user, f, e)}}
                       useNativeValidation={false}
                       label={f.name} 
                       type={f.type} />
            {#if f.invalid}
            <HelperText validationMsg>{f.error}</HelperText>
            {/if}
            <br />
          {/if}
        {/each}
        <br />
      </Content>
      <Actions>
        {#if have_account}
          <Button on:click={loginUser}>
            <Label>
              Login
            </Label>
            <Icon invalid class="material-icons">lock</Icon>
          </Button>
        {:else}
          <Button on:click={registerUser}>
            <Label>
              Register
            </Label>
          </Button>
        {/if}
        {#if $user.error}
        <p style="color: var(--mdc-theme-error);">
        {$user.error}
        </p>
        {/if}
      </Actions>
    </Card>
  {/if}
</div>


<script>
  import Checkbox from '@smui/checkbox';
  import FormField from '@smui/form-field';
  import Textfield from '@smui/textfield';
  import HelperText from '@smui/textfield/helper-text/index';
  import Button, {Icon, Label} from '@smui/button';
  import Card, {Content, Actions} from '@smui/card';
  import {user} from './auth';
  import {validateField, getValidData} from './utils';

  let have_account = true;
  let fields = [
      {
          name: "Nickname", key: "nickname", 
          value: "", register: false, 
          type: "nickname", invalid: false,
          error: ""
      },
      {
          name: "Password", key: "password", 
          value: "", register: false, 
          type: "password", invalid: false,
          error: ""
      },
      {
          name: "Firstname", key: "firstname", 
          value: "", register: true, 
          type: "text", invalid: false,
          error: ""
      },
      {
          name: "Lastname", key: "lastname", 
          value: "", register: true, 
          type: "text", invalid: false,
          error: ""
      },
      {
          name: "Email", key: "email", 
          value: "", register: true, 
          type: "email", invalid: false,
          error: ""
      },
      {name: "Phone", key: "phone", 
          value: "", register: true, 
          type: "phone", invalid: false,
          error: ""
      },
    ];

  function loginUser() {
    let data = getValidData(fields, user);
    if (!data) return;
    user.login(data);
  }
  function registerUser() {
    let data = getValidData(fields, user);
    if (!data) return;
    user.register(data);
  }
</script>
