// Composables
import {createRouter, createWebHistory} from 'vue-router'
import {chatSignRefresh, fetchUserInfo} from "@/api/api";
import {AppState} from "@/main";

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
    component: () => import('@/views/Signup.vue'),
    name: 'SignUp'
  },
  {
    path: '/login',
    component: () => import("@/views/Login.vue"),
    name: 'Login'
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
        component: () => import("@/components/docs/MDEditor.vue"),
      },
      {
        path: "files",
        name: "Files",
        component: () => import("@/components/files/FileList.vue")
      }
    ],
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

router.beforeEach(async (to) => {
    if (to.name == "Home" || to.name == "Login" || to.name == "SignUp") {
      return
    }
    if (AppState.sign_key == "empty") {
      await fetchUserInfo().then(response => {
        const {data} = response
        console.log(data)
        AppState.group_id = data.groupId
        AppState.user_id = data.userId
      }).then(async () => {
        const request = {sig: AppState.sign_key}
        AppState.sign_key = "empty"
        await chatSignRefresh(request).then(
          response => {
            const {data} = response
            if (!data.ok) {
              AppState.sign_key = data.sig
            }
          }
        )
      }).catch(error => {
        router.replace({path: "/"})
        console.log(error)
      })
    }
  }
)

export default router
