import { createSignal, lazy } from "solid-js";
const SideBar = lazy(() => import("./components/SideBar"))
const Header = lazy(() => import("./components/Header"))
const Main = lazy(() => import("./components/Main"))
const Footer = lazy(() => import("./components/Footer"))
function App() {
  return (
    <div class="flex flex-row flex-nowrap w-screen h-screen overflow-hidden">
      {/* sidebar */}
      <div class="bg-gray-500 w-16">
        <SideBar></SideBar>
      </div>
      {/* right */}
      <div class="flex flex-col w-full h-full">
        {/* header */}
        <div class="bg-yellow-200 h-14 w-full">
          <Header></Header>
        </div>
        {/* main */}
        <div class="bg-orange-300 w-full flex-grow overflow-auto">
          <Main></Main>
        </div>
        {/* footer */}
        <div class="bg-purple-400 h-14 w-full">
          <Footer></Footer>
        </div>
      </div>
    </div>
  );
}

export default App;
