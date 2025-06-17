import Topbar from "../components/Topbar";

const Home = () => {
  return (
    <div className="h-dvh w-full flex">
      <aside className="h-full bg-blue-500 p-4">
        <h2>sidebar</h2>
        <nav></nav>
      </aside>
      <div className="w-full flex flex-col">
        <Topbar />
        <main className="flex-1 bg-gray-100 p-4">
          <p>Main content area</p>
        </main>
      </div>
    </div>
  );
};

export default Home;
