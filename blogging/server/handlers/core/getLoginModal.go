package core

import "net/http"

var LoginModal HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// Endpoint: GET /LoginModal
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
    <div x-data={userField:"", passField:""} class="fixed top-0 left-0 bottom-0 right-0 bg-black bg-op-50 flex flex-col items-center justify-center z-36"
      id="modal" _="on closeModal add .closing then wait for animationend then remove me">
      <div class="modal-underlay absolute -z-1 top-0 left-0" _="on click trigger closeModal"></div>
      <div class="modal-content bg-[#584B77] w-96 h-max rounded-lg border border-2 border-black p-4">
        <div class="flex justify-around items-center">
          <h1 class="font-['Outfit'] text-white text-center align-baseline">Join Galactik<h1>
          <div _="on click trigger closeModal" class="bg-transparent rounded-[15px] border-2 border-[#27272a] flex justify-center items-center hover:bg-[#090909] transition ease-in-out transition-duration-200 cursor pointer w-max p-4 mx-2">
            <span class="font-normal text-[20px] font-['Outfit'] text-white align-baseline">
              X
            </span>
          </div>
        </div>
        <div class='px-4'>
          <span class="text-[20px] text-white font-['Outfit'] font-light">
            Username
          </span>
          <input x-data x-model="userField" type="text" id="username" class="flex h-9 w-82 rounded-md border border-input border-2 bg-transparent px-3 py-1 w-48 text-sm shadow-sm transition-colors text-white bg-op-50 font-['Outfit'] font-light" />
        </div>
        <div class='px-4 pb-4'>
          <span class="text-[20px] text-white font-['Outfit'] font-light">
            Password
          </span>
          <input x-data x-model="passField" type="password" id="password" class="flex h-9 w-82 rounded-md border border-input border-2 bg-transparent px-3 py-1 w-48 text-sm shadow-sm transition-colors text-white bg-op-50 font-['Outfit'] font-light" />
          <div @click:"console.log(passField, userField)" class="p-4 flex justify-center items-center bg-transparent hover:bg-[#090909] rounded-[15px] transition ease-in-out transition-duration-200 cursor pointer">
            <span class="font-normal text-2xl font-['Outfit'] text-white align-baseline ">
              Submit
            </span>
          </div>
        </div>
      </div>
    </div>
  `))
}
