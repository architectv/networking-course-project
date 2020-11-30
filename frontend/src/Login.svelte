
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
                       on:input={(e) => {onInput(i, e)}}
                       useNativeValidation={false}
                       label={f.name} 
                       type={f.type} />
            <HelperText validationMsg>{f.error}</HelperText>
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
  import {getUser} from './auth';

  let user = getUser();
  let have_account = true;
  let fields = [
      {
          name: "Nickname", key: "nickname", 
          value: "", register: false, 
          type: "nickname", invalid: false,
          updateInvalid: false, error: ""
      },
      {
          name: "Password", key: "password", 
          value: "", register: false, 
          type: "password", invalid: false,
          updateInvalid: false, error: ""
      },
      {
          name: "Firstname", key: "firstname", 
          value: "", register: true, 
          type: "text", invalid: false,
          updateInvalid: false, error: ""
      },
      {
          name: "Lastname", key: "lastname", 
          value: "", register: true, 
          type: "text", invalid: false,
          updateInvalid: false, error: ""
      },
      {
          name: "Email", key: "email", 
          value: "", register: true, 
          type: "email", invalid: false,
          updateInvalid: false, error: ""
      },
      {name: "Phone", key: "phone", 
          value: "", register: true, 
          type: "phone", invalid: false,
          updateInvalid: false, error: ""
      },
    ];

  function onInput(i, e) {
    if (e.srcElement) {
      let value = e.srcElement.value;
      let invalid = user.validate_prop(fields[i].key, value);
      if (invalid) {
        fields[i].invalid = true;
        fields[i].error = invalid;
        fields = fields;
      } else {
        fields[i].invalid = false;
      }
    }
  }

  function getData() {
    let data = {}
    for (let i in fields) {
      if (!have_account || !fields[i].register) {
        data[fields[i].key] = fields[i].value;
      }
    }
    let validation = user.validate(data);
    if (validation) {
      console.log("Login.getData: validation", validation);
      for (let i in fields) {
        if (fields[i].key in validation) {
          fields[i].invalid = true;
          fields[i].error = validation[fields[i].key];
        } else {
          fields[i].invalid = false;
        }
      }
      fields = fields;
      return null;
    }
    return data;
  }

  function loginUser() {
    let data = getData();
    if (!data) return;
    user.login(data);
  }
  function registerUser() {
    let data = getData();
    if (!data) return;
    user.register(data);
  }
</script>
