import axios from "axios";

export const login = (form) => {
  return axios.request({
    url: "/auth/login",
    method: "post",
    data: form,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
  })
}
export const logout = (form) => {
  return axios.request({
    url: "/auth/logout",
    method: "post",
    data: form,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
  })
}

export const translateBaidu = (form) => {
  return axios.request(
    {
      url: "/api/translate",
      method: "post",
      data: form,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    }
  )
}

export const chatSignRefresh = (form) => {
  return axios.request(
    {
      url: "/chat/refresh",
      method: "post",
      data: form,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    },
  )
}


export const deleteFiles = (form) => {
  return axios.request(
    {
      url: "/files/delete",
      method: "post",
      data: form,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    }
  )
}

export const listFiles = (query) => {
  return axios.request(
    {
      url: "/files/list?" + query,
      method: "get",
    }
  )
}

export const newDocs = (param) => {
  return axios.request({
    url: "/docs/new",
    method: "post",
    data: param,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

export const updateDocs = (param) => {
  return axios.request({
    url: "/docs/update",
    method: "post",
    data: param,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

export const deleteDocs = (param) => {
  return axios.request({
    url: "/docs/delete",
    method: "post",
    data: param,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

export const listDocs = (query) => {
  return axios.request(
    {
      url: "/docs/list?" + query,
      method: "get",
    }
  )
}


export const fetchUserInfo = () => {
  return axios.request(
    {
      url: "/auth/fetch-user-info",
      method: "get",
    }
  )
}

export const getDocs = (query) => {
  return axios.request(
    {
      url: "/docs/byid?" + query,
      method: "get",
    }
  )
}
