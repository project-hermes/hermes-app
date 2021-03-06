import firebase from '~/firebase';
const storageRef = firebase.storage().ref();

export default {
    namespaced: true,
    state: {
        status: [],
        progress: []
    },
    mutations: {
        setStatus(state, {status, index}) {
            state.status[index] = status;
            state.status = [...state.status];
        },
        resetState(state) {
            state.status = [];
            state.progress = [];
        },
        setProgress(state, {progress, index}) {
            state.progress[index] = progress;
            state.progress = [...state.progress];
        }
    },
    actions: {
        uploadFiles({commit, rootGetters}, {files}) {
            const promises = [];
            const user = rootGetters['auth/user'];
            const date = Date.now();
            files.forEach((file, index) => {
                commit('setStatus', {
                    status: 'uploading',
                    index
                });

                const diveRef = storageRef.child(
                    process.env.NODE_ENV !== 'production'
                        ? `TEST-raw-dive-files/${user.uid}-${date + index}`
                        : `raw-dive-files/${user.uid}-${date + index}`
                );
                const uploadTask = diveRef.put(file);
                promises.push(uploadTask);

                uploadTask.on(
                    'state_changed',
                    snapshot => {
                        const progress =
                            (snapshot.bytesTransferred / snapshot.totalBytes ||
                                0) * 100;
                        commit('setProgress', {
                            index,
                            progress
                        });
                    },
                    () => {
                        // error
                        commit('setStatus', {
                            status: 'error',
                            index
                        });
                    },
                    () => {
                        commit('setStatus', {
                            status: 'resolved',
                            index
                        });
                    }
                );
            });

            return Promise.all(promises);
        }
    },
    getters: {
        progress(state) {
            return state.progress;
        },
        status(state) {
            return state.status;
        }
    }
};
