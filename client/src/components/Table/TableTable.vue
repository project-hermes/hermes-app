<template>
  <table>
    <thead>
      <tr>
        <th
          v-for="column in columns"
          :key="column.key"
          scope="col">{{ column.label }}</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="row in rows"
        :key="row[idKey]">
        <td 
          v-for="column in columns"
          :key="column.key"
          scope="col">{{ row[column.key] }}</td>
      </tr>
    </tbody>
  </table>
</template>
<script>
export default {
    props: {
        idKey: {
            type: String
        },
        rows: {
            type: Array
        },
        columns: {
            type: Array
        }
    }
};
</script>
