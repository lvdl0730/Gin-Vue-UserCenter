import axios from 'axios'

//创建axios实例
const request = axios.create({
    baseURL: 'http://localhost:8080/',
    timeout: 5000,
})

//登录接口
export function login(data:{username:string;password:string}) {
    return request.post('/login',data);
}

//注册接口
export function register(data:{username:string;password:string}) {
    return request.post('/register',data);
}