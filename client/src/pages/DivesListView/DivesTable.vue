<template>
    <DumbTable
        :columns="columns"
        :rows="rows"
        :sort-column="sortColumn"
        id-key="id"
        class="rounded-sm"
        @sort="onSort"
    />
</template>
<script>
import {DumbTable} from '~/components';
import times from 'lodash/times';
import sortBy from 'lodash/sortBy';
export default {
    components: {
        DumbTable
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
            return this.rawRows.map(row => {
                return row;
            });
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
