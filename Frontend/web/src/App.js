import "./App.css";
import { UserTable } from "./components/userTable";

function App() {
  return (
    <div className="flex h-screen bg-gray-100">
      <aside className="w-64 bg-[#2C3E50] text-white">
        <nav className="mt-10">
          <ul>
            <li
              className={`pl-6 py-2 hover:bg-gray-700 
                         ${window.innerWidth <= 768 ? 'hidden md:block' : ''}`} // Responsive visibility for smaller screens
            >
              <LayoutDashboardIcon className="inline-block w-6 h-6 mr-2" />
              Dashboard{"\n                  "}
            </li>
            <li className="pl-6 py-2 bg-gray-700 m">
              <UsersIcon className="inline-block w-6 h-6 mr-2" />
              Users{"\n                  "}
            </li>
          </ul>
        </nav>
      </aside>
      <main className="flex-1">
        <header className="flex justify-center items-center p-6 bg-white">
          <h1 className="text-black text-bold text-xl font-semibold">Users</h1>
        </header>

        <div className="px-6 py-4">
          <div className="bg-white rounded shadow">
            <UserTable />
          </div>
        </div>
      </main>
    </div>
  );
}
function LayoutDashboardIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect width="7" height="9" x="3" y="3" rx="1" />
      <rect width="7" height="5" x="14" y="3" rx="1" />
      <rect width="7" height="9" x="14" y="12" rx="1" />
      <rect width="7" height="5" x="3" y="16" rx="1" />
    </svg>
  );
}

function UsersIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" />
      <circle cx="9" cy="7" r="4" />
      <path d="M22 21v-2a4 4 0 0 0-3-3.87" />
      <path d="M16 3.13a4 4 0 0 1 0 7.75" />
    </svg>
  );
}
export default App;
