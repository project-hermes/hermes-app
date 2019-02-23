<template>
  <Table
    :columns="columns"
    :rows="rows"
    id-key="id"
    class="rounded-sm"
  />

  <!-- <TableTable
    :columns="columns"
    :rows="rows"
    id-key="id"
  /> -->
</template>
<script>
import Table from './Table.vue';
import TableTable from './TableTable.vue';
export default {
    components: {
        Table,
        TableTable
    },
    data() {
        return {
            columns: [
                {key: 'date', label: 'Date'},
                {key: 'location', label: 'Location'},
                {key: 'country', label: 'Country'},
                {key: 'diver', label: 'Diver'},
                {key: 'computer', label: 'Computer'}
            ],
            rows: [
                {
                    id: '123',
                    date: 123,
                    location: 'gggg gggg ggg',
                    country: 'Greece',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '234',
                    date: 345094305,
                    location: 'somewhere',
                    country: 'merica',
                    diver: 'Bob',
                    computer: 'Scubutt'
                },
                {
                    id: '345',
                    date: 2323423,
                    location: 'a place',
                    country: 'Sealand',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '1231',
                    date: 123,
                    location: 'gggg gggg ggg',
                    country: 'Greece',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '2341',
                    date: 345094305,
                    location: 'somewhere',
                    country: 'merica',
                    diver: 'Bob',
                    computer: 'Scubutt'
                },
                {
                    id: '3451',
                    date: 2323423,
                    location: 'a place',
                    country: 'Sealand',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '1232',
                    date: 123,
                    location: 'gggg gggg ggg',
                    country: 'Greece',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '2342',
                    date: 345094305,
                    location: 'somewhere',
                    country: 'merica',
                    diver: 'Bob',
                    computer: 'Scubutt'
                },
                {
                    id: '3452',
                    date: 2323423,
                    location: 'a place',
                    country: 'Sealand',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '1233',
                    date: 123,
                    location: 'gggg gggg ggg',
                    country: 'Greece',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                },
                {
                    id: '2343',
                    date: 345094305,
                    location: 'somewhere',
                    country: 'merica',
                    diver: 'Bob',
                    computer: 'Scubutt'
                },
                {
                    id: '3453',
                    date: 2323423,
                    location: 'a place',
                    country: 'Sealand',
                    diver: 'Bob',
                    computer: 'Scubadooba'
                }
            ]
        };
    }
};
</script>
