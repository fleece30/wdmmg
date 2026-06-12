import { SidebarTrigger } from "./ui/sidebar"

const Navbar = () => {
  return (
    <nav className="flex items-center justify-between p-4">
      <SidebarTrigger />
      <div className="flex">Yello</div>
    </nav>
  )
}

export default Navbar
