
<pre>
{JSON.stringify(fields)}
</pre>

<br/>

<Fields bind:fields />

<script>
  import Fields from './Fields.svelte';
  let fields = [
      {
          name: "Title",
          value: "text",
          tag: "textfield",
      },
      {
          name: "Will be visible",
          value: false,
          tag: "checkbox",
          onChanged: (value) => {
            console.log(value);
            fields[0].invalid = true;
            fields = [...fields];
          }
      },
      {
          name: "Will be hidden",
          value: false,
          tag: "checkbox",
          filtrate: true
      }
  ];

</script>
