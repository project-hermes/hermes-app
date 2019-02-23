<template>
  <div>
    <div
    class="p-2 bg-grey-lightest border-b solid border-grey-lighter text-black font-bold whitespace-no-wrap overflow-x-auto">
      {{ column.label }}
    </div>
    <div
      v-for="row in rows"
      :key="row[idKey]"
      class="p-2 border-b solid border-grey-lighter text-black whitespace-no-wrap overflow-x-auto">
      {{ row[column.key] }}
    </div>
  </div>
</template>
<script>
export default {
    props: {
        column: {
            type: Object
        },
        rows: {
            type: Array
        },
        idKey: {
            type: String
        }
    }
};
</script>
