import axios from 'axios';


console.log(process.env.NODE_ENV)


var instance = axios.create();

instance.defaults.baseURL = 'http://127.0.0.1'
instance.defaults.timeout = 1000



instance.interceptors.request.use(
    config =>{
        console.log(config)
        return config
    },
    error =>{
        console.log(error)
    }
)

instance.interceptors.response.use(
    response =>{
        console.log(response)
        if (response.status == 200){
            return Promise.resolve(response)
        } else{
            return Promise.reject(response)
        }
    },
    error => {

        const { response } = error;
        if (response) {
            return Promise.reject(response);
        } else {
            // store.commit('changeNetwork', false);
        }

        // console.log(error)
    }
)



export default instance;