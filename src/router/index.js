// Composables
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/homepage/Default.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        // route level code-splitting
        // this generates a separate chunk (Home-[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import('@/views/Home.vue'),
      },
    ],
  },
  {
    path: '/signup',
    component: () => import('@/views/Signup.vue')
  },
  {
    path: '/login',
    component: () => import("@/views/Login.vue")
  },
  {
    path: '/workspace',
    component: () => import("@/layouts/workspace/Default.vue"),
    children: [
      {
        path: "",
        name: "WorkspaceHome",
        component: () => import("@/views/WorkspaceHome.vue")
      },
      {
        path: "userinfo",
        name: "UserInfo",
        component: () => import("@/components/UserInfo.vue")
      },
      {
        path: "docs",
        name: "Docs",
        component: () => import("@/components/docs/DocsList.vue"),
      },
      {
        path: "docs/editor",
        name: "Editor",
        component: () => import("@/components/docs/MDEditor.vue")
      },
      {
        path: "docs/view/:id",
        name: "Preview",
        component: () => import("@/components/docs/DocsPreview.vue")
      },
      {
        path: "docs/edit/:id",
        name: "Edit",
        component: () => import("@/components/docs/MDEditor.vue")
      },
      {
        path: "files",
        name: "Files",
        component: () => import("@/components/files/FileList.vue")
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
