<template>
    <div class="">
        <div
            @click="onSort"
            class="h-10 px-2 flex flex-row items-center justify-between bg-grey-lightest border-b solid border-grey-lighter text-black font-bold whitespace-no-wrap overflow-x-auto cursor-pointer hover:bg-grey"
        >
            <span class="ml-1">{{ column.label }}</span>
            <span class="mr-1 md:mr-4 h-6">
                <ChevronUpIcon
                    v-if="sortOrder === 'ascending'"
                    class="text-blue"
                />
                <ChevronDownIcon
                    v-else-if="sortOrder === 'descending'"
                    class="text-blue"
                />
            </span>
        </div>
        <RecycleScroller :items="rows" :item-size="40" 
:key-field="idKey">
            <template v-slot="{item: row}">
                <div
                    class="h-10 p-2 flex items-center border-b solid border-grey-lighter text-black whitespace-no-wrap overflow-x-auto"
                >
                    <a
                        v-if="column.linkProp"
                        :href="row[column.linkProp]"
                        class="ml-1 no-underline text-blue hover:underline"
                    >
                        {{ row[column.key] }}
                    </a>
                    <span v-else class="ml-1">
                        {{ row[column.key] }}
                    </span>
                </div>
            </template>
        </RecycleScroller>
    </div>
</template>
<script>
import ChevronUpIcon from 'vue-feather-icons/icons/ChevronUpIcon';
import ChevronDownIcon from 'vue-feather-icons/icons/ChevronDownIcon';

export default {
    components: {
        ChevronUpIcon,
        ChevronDownIcon
    },
    props: {
        column: {
            type: Object
        },
        rows: {
            type: Array
        },
        idKey: {
            type: String
        },
        sortColumn: {
            type: Object,
            default: () => ({})
        }
    },
    computed: {
        sortOrder() {
            if (this.sortColumn.key === this.column.key) {
                return this.sortColumn.order;
            }
        }
    },
    methods: {
        onSort() {
            this.$emit('sort', {
                key: this.column.key,
                order:
                    this.sortOrder === 'ascending' ? 'descending' : 'ascending'
            });
        }
    }
};
</script>
