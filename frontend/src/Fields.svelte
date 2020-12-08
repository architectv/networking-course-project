<p>
FIELDS: {JSON.stringify(fields)}
</p>


{#each fields as f, i}
  {#if !f.invisible}
    {#if f.tag == 'textfield'}
      <Textfield bind:invalid={f.invalid}
                 bind:value={f.value} 
                 on:input={(e) => {f.onInput(i, e)}}
                 useNativeValidation={false}
                 label={f.name} 
                 type={f.type} />
      {#if f.invalid}
        <HelperText validationMsg>{f.error}</HelperText>
      {/if}
    {:else if f.tag == 'checkbox'}
      <FormField>
        <Checkbox bind:checked={f.value} on:change={f.onChanged} />
      <Label>{f.name}</Label>
      </FormField>
    {/if}
    <br />
  {/if}
{/each}

<script>
  import HelperText from '@smui/textfield/helper-text/index';
  import Button, {Icon, Label} from '@smui/button';
  import Checkbox from '@smui/checkbox';
  import FormField from '@smui/form-field';
  import Textfield from '@smui/textfield';
  export let fields = [];

</script>
