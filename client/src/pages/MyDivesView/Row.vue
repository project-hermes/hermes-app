<template>
  <div>
    <div
    class="px-2 py-3 inline-flex border-b solid border-black text-black flex-1 items-stretch "
      v-for="{key} in columns"
      :key="key">
      {{ row[key] }}
    </div>
  </div>
</template>
<script>
import Row from './Row.vue';
export default {
    components: {
        Row
    },
    props: {
        row: {
            type: Object
        },
        columns: {
            type: Array
        }
    }
};
</script>
