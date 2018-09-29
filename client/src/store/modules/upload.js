import firebase from '~/firebase';
const storageRef = firebase.storage().ref();

export default {
    namespaced: true,
    state: {
        status: '',
        fileName: '',
        progress: 0,
        file: null
    },
    mutations: {
        setStatus(state, status) {
            state.status = status;
        },
        resetStatus(state) {
            state.status = '';
        },
        setProgress(state, progress) {
            state.progress = progress;
        }
    },
    actions: {
        uploadFile({commit, rootGetters}, {file}) {
            commit('setStatus', 'uploading');
            const user = rootGetters['auth/user'];
            try {
                const diveRef = storageRef.child(
                    `raw-dive-files/${user.uid}-${Date.now()}`
                );
                const uploadTask = diveRef.put(file);

                uploadTask.on(
                    'state_changed',
                    snapshot => {
                        const progress =
                            (snapshot.bytesTransferred / snapshot.totalBytes ||
                                0) * 100;
                        commit('setProgress', progress);
                    },
                    () => {
                        // errrrrr
                        commit('setStatus', 'error');
                    },
                    () => {
                        // success
                        commit('setStatus', 'complete');
                    }
                );
            } catch (e) {
                commit('setStatus', 'error');
            }
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
