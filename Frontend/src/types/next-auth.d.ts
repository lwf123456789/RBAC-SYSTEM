import NextAuth from "next-auth"

declare module "next-auth" {
  interface Session {
    accessToken?: string
    user: {
      id: string
      name: string
      email: string
      phone: string
      department_id: number
      status: number
      last_login: string
      created_at: string
      updated_at: string
      roles: Array<{
        id: number
        name: string
        description: string
        created_at: string
        updated_at: string
      }>
      department: {
        id: number
        name: string
        parent_id: number | null
        icon: string
        created_at: string
        updated_at: string
      }
    }
  }

  interface User {
    token: string
    user: {
      id: number
      name: string
      email: string
      phone: string
      department_id: number
      status: number
      last_login: string
      created_at: string
      updated_at: string
      roles: Array<{
        id: number
        name: string
        description: string
        created_at: string
        updated_at: string
      }>
      department: {
        id: number
        name: string
        parent_id: number | null
        icon: string
        created_at: string
        updated_at: string
      }
    }
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    accessToken?: string
    user?: {
      id: number
      name: string
      email: string
      phone: string
      department_id: number
      status: number
      last_login: string
      created_at: string
      updated_at: string
      roles: Array<{
        id: number
        name: string
        description: string
        created_at: string
        updated_at: string
      }>
      department: {
        id: number
        name: string
        parent_id: number | null
        icon: string
        created_at: string
        updated_at: string
      }
    }
  }
}