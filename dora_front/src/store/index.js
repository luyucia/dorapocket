import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        data: [],
        form_data:{
        },
        form_comp: {},
        layout:[]
    },
    mutations: {
        setState(state, data) {
            state.data = data
        },
        addComp(state, data) {
            Object.assign(state.form_comp,data)
            console.log(state.form_comp)
        },
        setComp(state, data){
            console.log(data)
            state.form_comp.c0.label = 'sss'
        }
    },
    actions: {},
    modules: {}
})
