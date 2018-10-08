<template>
  <div
    class="relative"
    @drop="onDrop"
    @dragover="onDragOver"
    @dragleave="onDragLeave"
  >
    <input
      ref="input"
      class=""
      hidden
      aria-hidden="true"
      tabindex="-1"
      type="file"
      @change="onFileSelect">
    <div
      v-if="files.length === 0"
      :class="[isDraggingOver ? 'border-black' : 'border-grey']"
      class="absolute h-full w-full border border-dashed rounded flex flex-col items-center justify-center">
      <UploadIcon class="text-grey h-16 w-16" />
      <p class="py-3">
        Drag and drop a file here or
        <a
          class="hover:underline text-blue cursor-pointer"
          @click="openBrowse">click to upload</a>
      </p>
    </div>
    <div
      v-else
      class="w-full h-full">
      <FileItem
        v-for="(file, index) in files"
        :file="file"
        :index="index"
        :key="index"
        :progress="progress[index]"
        @remove="onRemove" />
      <div class="px-4 py-2 flex justify-end">
        <button
          class="bg-white hover:bg-grey-lighter text-black font-normal text-sm my-2 py-2 px-4 rounded"
          @click="openBrowse"
        >
          Upload more files
        </button>
        <button
          class="bg-blue hover:bg-blue-dark text-white font-normal text-sm my-2 py-2 px-4 rounded ml-1"
          @click="onConfirm">
          Submit
        </button>
      </div>
    </div>
  </div>
</template>
<script>
import UploadIcon from 'vue-feather-icons/icons/UploadIcon';
import FileItem from './FileItem.vue';
export default {
    components: {
        UploadIcon,
        FileItem
    },
    props: {
        progress: {
            type: Array,
            default: []
        }
    },
    data() {
        return {
            isDraggingOver: false,
            files: []
        };
    },
    methods: {
        openBrowse() {
            this.$refs.input.click();
        },
        onDrop(e) {
            e.preventDefault();

            if (e.dataTransfer.items) {
                // Use DataTransferItemList interface to access the file(s)
                for (var i = 0; i < e.dataTransfer.items.length; i++) {
                    // If dropped items aren't files, reject them
                    if (e.dataTransfer.items[i].kind === 'file') {
                        var file = e.dataTransfer.items[i].getAsFile();
                        this.files.push(file);
                    }
                }
            } else {
                // Use DataTransfer interface to access the file(s)
                this.files = this.files.concat(e.dataTransfer.files);
            }

            // Pass event to removeDragData for cleanup
            this.removeDragData(e);
            this.isDraggingOver = false;
        },
        removeDragData(e) {
            if (e.dataTransfer.items) {
                // Use DataTransferItemList interface to remove the drag data
                e.dataTransfer.items.clear();
            } else {
                // Use DataTransfer interface to remove the drag data
                e.dataTransfer.clearData();
            }
        },
        onDragOver(e) {
            e.preventDefault();
            this.isDraggingOver = true;
        },
        onDragLeave(e) {
            this.isDraggingOver = false;
        },
        onFileSelect(e) {
            const files = [];
            for (let i = 0; i < e.target.files.length; i++) {
                files.push(e.target.files.item(i));
            }
            this.files = this.files.concat(files);
        },
        onRemove(i) {
            this.files.splice(i, 1);
        },
        onConfirm(e) {
            this.$emit('confirm', this.files);
        }
    }
};
</script>
