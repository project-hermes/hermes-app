<template>
    <div class="flex border solid border-grey-lighter">
        <Column
            v-for="column in columns"
            :key="column.key"
            :column="column"
            :rows="rows"
            :id-key="idKey"
            class="flex-auto"
        />

        <!-- <div>
            <div
            class="p-2 inline-flex border-b solid border-black text-black flex-1 items-stretch "
              v-for="{key, label} in columns"
              :key="key">
              {{label}}
            </div>
        </div>
        <Row
            class="w-full"
            v-for="row in rows"
            :key="row[idKey]"
            :row="row"
            :columns="columns"
        /> -->
    </div>
</template>
<script>
    import Column from './Column.vue';
    export default {
        components: {
            Column
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
            }
        }
    }
</script>
