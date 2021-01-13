import http from './http'



const project ={

    List() {
        return http.get('/project/list')
    },
    Delete(id) {
        return http.post('/project/delete',{id:id})
    },
    Create(data) {
        return http.post('/project/create',data)
    },
    Update(data) {
        return http.post('/project/edit',data)
    },
}

export default project;


