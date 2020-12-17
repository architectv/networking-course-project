<Dialog bind:this={newDialog}>
  <Title>{title}</Title>
  <Content>
    {subtitle}
        <br />
        {#if fields}
        {#each fields as f, i}
          {#if f.type === "checkbox"}
          <FormField>
            <Checkbox bind:value={f.value} />
            <span slot="label">{f.name}</span>
          </FormField>
          {:else}
            <Textfield bind:invalid={f.invalid}
                       bind:value={f.value} 
                       on:input={(e) => {validateField(objs, f, e)}}
                       useNativeValidation={false}
                       label={f.name} 
                       textarea={f.long}
                       style="min-width: 25%;"
                       type={f.type} />
            {#if f.invalid}
            <HelperText validationMsg>{f.error}</HelperText>
            {/if}
          {/if}
            <br />
        {/each}
        {/if}
        {#if description}
          {@html marked(description)}
        {/if}
        <br />
  </Content>
  <Actions>
    <Button on:click={_onCreate}>
    {#if isUpdate}
      <Label>Change</Label>
    {:else}
      <Label>Create</Label>
    {/if}
    </Button>
    <Button on:click={() => {}}>
      <Label>Cancel</Label>
    </Button>
  </Actions>
</Dialog>

<script>
  import marked from 'marked';
  import HelperText from '@smui/textfield/helper-text/index';
  import FormField from '@smui/form-field'; 
  import Textfield from '@smui/textfield';
  import {validateField, getValidData} from '../utils';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Button, {Icon, Label} from '@smui/button';
  import Checkbox from '@smui/checkbox';
  
  let newDialog;
  let fields;
  let objs;
  let desc_id;
  let title;
  let subtitle;
  let isUpdate;
  $: description = desc_id != undefined && fields[desc_id].value;
  
  export let onCreate;
  function _onCreate() {
    if (onCreate) onCreate(fields);
  }
  
  export function open(objs_, fields_, desc_id_, onCreate_, title_, 
                       subtitle_, isUpdate_=false, ...args) {
    console.log(objs_, fields_, desc_id_, onCreate_, title_, subtitle_);
    fields = fields_;
    objs = objs_;
    desc_id = desc_id_;
    onCreate = onCreate_;
    title = title_;
    subtitle = subtitle_;
    isUpdate = isUpdate_;
    newDialog.open(...args)
  }
</script>