
<div style="background-color: {toColor(label.color)}; 
               border-radius: 10px; padding-left: 10px;">
  <div style="color: {toColor(label.color)}; filter: invert(100%); width: max-content;">
    <Button style="min-width: unset; padding: 0; margin: 0; height: 18px">
      <Label style="color: {toColor(label.color)}; font-size: 0.75rem">
        {label.name || "Unnamed"}
      </Label>
    </Button>
    {#if onDelete}
    <Button style="min-width: unset; padding: 0; margin: 0; height: 18px" on:click={onDelete}>
      <Icon style="color: {toColor(label.color)}; font-size: 14px; height: unset; width: unset;" 
            class="material-icons">delete_outline</Icon>
    </Button>
    {:else}
    <span style="padding-right: 10px;"></span>
    {/if}
  </div>
</div>
                        
<script>
  import Button, {Icon, Label} from '@smui/button';
  import {toColor} from '../utils';
  export let label;
  export let onDelete = undefined;
</script>