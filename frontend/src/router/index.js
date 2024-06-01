import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/Index.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login,
  },
  {
    path: '/Register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/index',
    name: 'Index',
    component: Index,
    beforeEnter: (to, from, next) => {
      // 尝试从 sessionStorage 中获取 userID
      const userID = sessionStorage.getItem('userID');

      if (!userID) {
        alert('The page you are visiting does not exist.');
        next('/'); // 重定向到登录页面
      } else {
        next(); // 允许导航
      }
    },
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})

export default router
