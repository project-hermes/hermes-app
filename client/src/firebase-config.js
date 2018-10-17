const config =
    process.env.NODE_ENV !== 'production'
        ? {
              apiKey: 'AIzaSyDl1MUKpDPJlLcXFs3MpcHWog8qZE-3c3s',
              authDomain: 'project-hermes-staging.firebaseapp.com',
              databaseURL: 'https://project-hermes-staging.firebaseio.com',
              projectId: 'project-hermes-staging',
              storageBucket: 'project-hermes-staging.appspot.com',
              messagingSenderId: '297733341091'
          }
        : {
              apiKey: 'AIzaSyDTsjfwJsoAFG2xaTXquPjawyjpaypGHiU',
              authDomain: 'project-hermes-6c5cb.firebaseapp.com',
              databaseURL: 'https://project-hermes-6c5cb.firebaseio.com',
              projectId: 'project-hermes-6c5cb',
              storageBucket: 'project-hermes-6c5cb.appspot.com',
              messagingSenderId: '185837078935'
          };

export default config;
