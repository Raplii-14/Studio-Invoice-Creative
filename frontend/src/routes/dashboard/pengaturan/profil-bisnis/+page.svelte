<script>
    let namaBisnis = $state("Kilat Studio");
    let emailBisnis = $state("hello@kilatstudio.id");
    let teleponBisnis = $state("+62 21 5555 0123");
    let alamatBisnis = $state("Jl. Sudirman No. 45, Jakarta");

    let showEditModal = $state(false);

    let tempNama = $state(namaBisnis);
    let tempEmail = $state(emailBisnis);
    let tempTelepon = $state(teleponBisnis);
    let tempAlamat = $state(alamatBisnis);

    function openModal() {
        tempNama = namaBisnis;
        tempEmail = emailBisnis;
        tempTelepon = teleponBisnis;
        tempAlamat = alamatBisnis;
        showEditModal = true;
    }

    function handleSave(e) {
        e.preventDefault();
        namaBisnis = tempNama;
        emailBisnis = tempEmail;
        teleponBisnis = tempTelepon;
        alamatBisnis = tempAlamat;
        showEditModal = false;
    }
</script>

<div class="space-y-6 max-w-3xl">
    <!-- Header Bagian -->
    <div
        class="border-b border-slate-100 pb-4 flex items-center justify-between"
    >
        <div>
            <h3 class="text-xl font-black text-slate-900 tracking-tight">
                Profil Bisnis
            </h3>
            <p class="text-xs text-slate-400 font-medium mt-0.5">
                Informasi ini tampil di kop faktur Anda
            </p>
        </div>
        <button
            onclick={openModal}
            class="px-4 py-2.5 rounded-2xl font-bold text-xs text-white bg-slate-900 hover:bg-slate-800 shadow-md shadow-slate-950/10 transition-all hover:-translate-y-0.5 flex items-center gap-2"
        >
            <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2.5"
                    d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"
                />
            </svg>
            Ubah Detail
        </button>
    </div>

    <!-- Detail Informasi Profil Bisnis -->
    <div
        class="bg-white rounded-3xl border border-slate-200/80 p-8 space-y-6 shadow-sm"
    >
        <!-- Baris Nama Bisnis & Email -->
        <div
            class="grid grid-cols-1 sm:grid-cols-2 gap-6 pb-6 border-b border-slate-100"
        >
            <div class="flex items-center gap-4">
                <div
                    class="w-10 h-10 rounded-2xl bg-slate-50 border border-slate-200/80 flex items-center justify-center text-slate-600 shrink-0"
                >
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                        />
                    </svg>
                </div>
                <div>
                    <p
                        class="text-[11px] font-bold text-slate-400 uppercase tracking-wider"
                    >
                        Nama Bisnis
                    </p>
                    <p class="text-sm font-semibold text-slate-800 mt-0.5">
                        {namaBisnis}
                    </p>
                </div>
            </div>

            <div class="flex items-center gap-4">
                <div
                    class="w-10 h-10 rounded-2xl bg-slate-50 border border-slate-200/80 flex items-center justify-center text-slate-600 shrink-0"
                >
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                        />
                    </svg>
                </div>
                <div>
                    <p
                        class="text-[11px] font-bold text-slate-400 uppercase tracking-wider"
                    >
                        Email
                    </p>
                    <p class="text-sm font-semibold text-slate-800 mt-0.5">
                        {emailBisnis}
                    </p>
                </div>
            </div>
        </div>

        <!-- Baris Telepon & Alamat -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <div class="flex items-center gap-4">
                <div
                    class="w-10 h-10 rounded-2xl bg-slate-50 border border-slate-200/80 flex items-center justify-center text-slate-600 shrink-0"
                >
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
                        />
                    </svg>
                </div>
                <div>
                    <p
                        class="text-[11px] font-bold text-slate-400 uppercase tracking-wider"
                    >
                        Telepon
                    </p>
                    <p class="text-sm font-semibold text-slate-800 mt-0.5">
                        {teleponBisnis}
                    </p>
                </div>
            </div>

            <div class="flex items-center gap-4">
                <div
                    class="w-10 h-10 rounded-2xl bg-slate-50 border border-slate-200/80 flex items-center justify-center text-slate-600 shrink-0"
                >
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                        />
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
                        />
                    </svg>
                </div>
                <div>
                    <p
                        class="text-[11px] font-bold text-slate-400 uppercase tracking-wider"
                    >
                        Alamat
                    </p>
                    <p class="text-sm font-semibold text-slate-800 mt-0.5">
                        {alamatBisnis}
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Modal / Popup Ubah Profil Bisnis -->
{#if showEditModal}
    <div
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6 font-sans"
    >
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div
            class="absolute inset-0 bg-slate-900/40 backdrop-blur-sm transition-opacity"
            onclick={() => (showEditModal = false)}
        ></div>

        <div
            class="relative w-full max-w-lg bg-white rounded-3xl shadow-2xl border border-slate-100 overflow-hidden transform transition-all flex flex-col"
        >
            <div
                class="px-8 py-6 border-b border-slate-100 flex items-center justify-between shrink-0 bg-white"
            >
                <div>
                    <h3
                        class="text-xl font-black text-slate-900 tracking-tight"
                    >
                        Ubah Profil Bisnis
                    </h3>
                    <p class="text-xs text-slate-400 font-medium mt-0.5">
                        Perbarui informasi kop faktur Anda di sini.
                    </p>
                </div>
                <button
                    onclick={() => (showEditModal = false)}
                    class="text-slate-400 hover:text-slate-700 transition-colors p-2 rounded-2xl hover:bg-slate-100"
                >
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d="M6 18L18 6M6 6l12 12"
                        /></svg
                    >
                </button>
            </div>

            <form onsubmit={handleSave} class="px-8 py-6 space-y-4">
                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Nama Bisnis</label
                    >
                    <input
                        type="text"
                        bind:value={tempNama}
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Email</label
                    >
                    <input
                        type="email"
                        bind:value={tempEmail}
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Telepon</label
                    >
                    <input
                        type="text"
                        bind:value={tempTelepon}
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Alamat</label
                    >
                    <input
                        type="text"
                        bind:value={tempAlamat}
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div
                    class="pt-4 border-t border-slate-100 flex justify-end gap-3"
                >
                    <button
                        type="button"
                        onclick={() => (showEditModal = false)}
                        class="px-6 py-2.5 rounded-xl font-bold text-sm text-slate-600 bg-white border border-slate-200 hover:bg-slate-50 transition-colors"
                    >
                        Batal
                    </button>
                    <button
                        type="submit"
                        class="px-6 py-2.5 rounded-xl font-bold text-sm text-white bg-slate-900 hover:bg-slate-800 shadow-lg shadow-slate-900/20 transition-all hover:-translate-y-0.5"
                    >
                        Simpan Perubahan
                    </button>
                </div>
            </form>
        </div>
    </div>
{/if}
