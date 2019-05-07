<template>
    <div class="flex border-t border-l border-r solid border-grey-lighter overflow-x-auto">
        <DumbColumn
            v-for="column in columns"
            :key="column.key"
            :column="column"
            :rows="rows"
            :id-key="idKey"
            :sort-column="sortColumn"
            class="flex-auto"
            @sort="onSort"
        />
    </div>
</template>
<script>
    import DumbColumn from './DumbColumn.vue';
    export default {
        components: {
            DumbColumn
        },
        props: {
            idKey: {
                type: String
            },
            rows: {
                type: Array
            },
            columns: {
                type: Array
            },
            sortColumn: {
                type: Object
            }
        },
        methods: {
            onSort (sortColumn) {
                this.$emit('sort', sortColumn);
            }
        }
    }
</script>
