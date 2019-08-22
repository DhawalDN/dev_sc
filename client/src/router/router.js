import Vue from 'vue'
import Router from 'vue-router'
import register from '@/components/register'
import students from '@/components/students'


Vue.use(Router)


export default new Router({
    routes: [
      {
        path: '/',
        redirect:'/register'
      },
      {
        path: '/register',
        name: 'register',
        component: register
      },
      {
        path: '/students',
        name: 'students',
        component: students
        
      }
    ]
  })