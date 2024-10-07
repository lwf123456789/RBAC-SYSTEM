import { useSession } from "next-auth/react"
import { useRouter } from "next/navigation"
import { Spin } from 'antd'

export function withAuth<P extends object>(WrappedComponent: React.ComponentType<P>) {
  return function WithAuth(props: P) {
    const { data: session, status } = useSession()
    const router = useRouter()

    if (status === "loading") {
      return <div className="flex justify-center items-center h-screen"><Spin size="large" /></div>
    }

    if (status === "unauthenticated") {
      router.replace("/")
      return null
    }

    return <WrappedComponent {...props} />
  }
}