<template>
    <DumbTable
        v-if="rows.length > 0"
        :columns="columns"
        :rows="rows"
        :sort-column="sortColumn"
        id-key="id"
        class="rounded-sm"
        @sort="onSort"
    />
    <div v-else class="h-64 flex flex-col items-center justify-center">
        <h2 class="font-normal text-center">
            You don't have any dives recorded in Project Hermes yet.
        </h2>
        <span class="text-center text-grey-darker m-4"
            >Sync data from your Hermes sensor to get started.</span
        >
    </div>
</template>
<script>
import {DumbTable} from '~/components';
import times from 'lodash/times';
import sortBy from 'lodash/sortBy';
import isFinite from 'lodash/isFinite';
export default {
    components: {
        DumbTable
    },
    props: {
        limit: {
            type: Number
        }
    },
    data() {
        return {
            columns: [
                {key: 'date', label: 'Date & Time', linkProp: 'divePath'},
                {key: 'coordinates', label: 'Coordinates'},
                {key: 'country', label: 'Country'},
                {key: 'duration', label: 'Duration'},
                {key: 'maxDepth', label: 'Max Depth'},
                {key: 'maxTemp', label: 'Max Temp'},
                {key: 'minTemp', label: 'Min Temp'}
            ],
            rawRows: [],
            sortColumn: {}
        };
    },
    computed: {
        rows() {
            return isFinite(this.limit)
                ? this.rawRows.slice(0, this.limit)
                : this.rawRows;
        }
    },
    mounted() {
        this.rawRows = times(100, i => {
            return {
                id: `${i}`,
                date: Date.now(),
                coordinates: [i, i],
                country: 'foo',
                duration: 1000 + i,
                maxDepth: 200 + i,
                maxTemp: 25 + i,
                minTemp: 25 - i,
                divePath: `/dives/${i}`
            };
        });

        this.onSort({
            key: 'duration',
            order: 'ascending'
        });
    },
    methods: {
        onSort(sortColumn) {
            this.sortColumn = sortColumn;
            const rawRows = sortBy(this.rawRows, [sortColumn.key]);
            this.rawRows =
                sortColumn.order === 'descending' ? rawRows.reverse() : rawRows;
        }
    }
};
</script>
