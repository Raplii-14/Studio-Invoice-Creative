<script>
    let activeTheme = $state("minimal");
    let showPreview = $state(false);
    let previewTheme = $state("minimal");

    function openPreview(theme) {
        previewTheme = theme;
        showPreview = true;
    }

    function applyTheme() {
        activeTheme = previewTheme;
        showPreview = false;
    }
</script>

<div class="space-y-6 max-w-4xl">
    <!-- Header Bagian -->
    <div
        class="border-b border-slate-100 pb-4 flex items-center justify-between"
    >
        <div>
            <h3 class="text-xl font-black text-slate-900 tracking-tight">
                Tampilan
            </h3>
            <p class="text-xs text-slate-400 font-medium mt-0.5">
                Sesuaikan preferensi visual dan tema faktur Anda
            </p>
        </div>
    </div>

    <!-- Kontainer Utama Pilihan Tema -->
    <div
        class="bg-white rounded-3xl border border-slate-200/80 p-8 space-y-6 shadow-sm"
    >
        <div>
            <h4 class="text-sm font-bold text-slate-900">
                Tema Faktur Default
            </h4>
            <p class="text-[11px] font-medium text-slate-400 mt-0.5">
                Klik salah satu tema di bawah ini untuk melihat *preview*
                (pratinjau) faktur secara penuh.
            </p>
        </div>

        <!-- Grid Pilihan Tema -->
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 pt-2">
            <!-- Opsi 1: Minimal Putih -->
            <button
                onclick={() => openPreview("minimal")}
                class="text-left group relative p-5 rounded-2xl border-2 transition-all duration-200 {activeTheme ===
                'minimal'
                    ? 'border-slate-900 bg-slate-50/50 shadow-md'
                    : 'border-slate-100 hover:border-slate-300 bg-white'}"
            >
                {#if activeTheme === "minimal"}
                    <div
                        class="absolute top-3 right-3 w-5 h-5 bg-slate-900 rounded-full flex items-center justify-center shadow-sm z-10"
                    >
                        <svg
                            class="w-3 h-3 text-white"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            ><path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="3"
                                d="M5 13l4 4L19 7"
                            /></svg
                        >
                    </div>
                {/if}
                <div
                    class="w-full h-32 bg-white rounded-xl border border-slate-200 shadow-sm mb-4 p-4 flex flex-col justify-between transition-transform duration-300 group-hover:scale-[1.03]"
                >
                    <div class="w-1/3 h-2 bg-slate-200 rounded-full"></div>
                    <div class="space-y-2">
                        <div
                            class="w-full h-1.5 bg-slate-100 rounded-full"
                        ></div>
                        <div
                            class="w-2/3 h-1.5 bg-slate-100 rounded-full"
                        ></div>
                    </div>
                </div>
                <div>
                    <p class="text-sm font-bold text-slate-900">
                        Minimal Putih
                    </p>
                    <p class="text-[10px] text-slate-500 font-medium mt-0.5">
                        Bersih & profesional
                    </p>
                </div>
            </button>

            <!-- Opsi 2: Gelap Elegan -->
            <button
                onclick={() => openPreview("gelap")}
                class="text-left group relative p-5 rounded-2xl border-2 transition-all duration-200 {activeTheme ===
                'gelap'
                    ? 'border-slate-900 bg-slate-50/50 shadow-md'
                    : 'border-slate-100 hover:border-slate-300 bg-white'}"
            >
                {#if activeTheme === "gelap"}
                    <div
                        class="absolute top-3 right-3 w-5 h-5 bg-slate-900 rounded-full flex items-center justify-center shadow-sm z-10"
                    >
                        <svg
                            class="w-3 h-3 text-white"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            ><path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="3"
                                d="M5 13l4 4L19 7"
                            /></svg
                        >
                    </div>
                {/if}
                <div
                    class="w-full h-32 bg-slate-900 rounded-xl border border-slate-800 shadow-sm mb-4 p-4 flex flex-col justify-between transition-transform duration-300 group-hover:scale-[1.03]"
                >
                    <div class="w-1/3 h-2 bg-slate-700 rounded-full"></div>
                    <div class="space-y-2">
                        <div
                            class="w-full h-1.5 bg-slate-800 rounded-full"
                        ></div>
                        <div
                            class="w-2/3 h-1.5 bg-slate-800 rounded-full"
                        ></div>
                    </div>
                </div>
                <div>
                    <p class="text-sm font-bold text-slate-900">Gelap Elegan</p>
                    <p class="text-[10px] text-slate-500 font-medium mt-0.5">
                        Modern & mewah
                    </p>
                </div>
            </button>

            <!-- Opsi 3: Hangat Krem -->
            <button
                onclick={() => openPreview("krem")}
                class="text-left group relative p-5 rounded-2xl border-2 transition-all duration-200 {activeTheme ===
                'krem'
                    ? 'border-slate-900 bg-slate-50/50 shadow-md'
                    : 'border-slate-100 hover:border-slate-300 bg-white'}"
            >
                {#if activeTheme === "krem"}
                    <div
                        class="absolute top-3 right-3 w-5 h-5 bg-slate-900 rounded-full flex items-center justify-center shadow-sm z-10"
                    >
                        <svg
                            class="w-3 h-3 text-white"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            ><path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="3"
                                d="M5 13l4 4L19 7"
                            /></svg
                        >
                    </div>
                {/if}
                <div
                    class="w-full h-32 bg-[#FDFBF7] rounded-xl border border-amber-200/60 shadow-sm mb-4 p-4 flex flex-col justify-between transition-transform duration-300 group-hover:scale-[1.03]"
                >
                    <div class="w-1/3 h-2 bg-amber-200/80 rounded-full"></div>
                    <div class="space-y-2">
                        <div
                            class="w-full h-1.5 bg-amber-100/80 rounded-full"
                        ></div>
                        <div
                            class="w-2/3 h-1.5 bg-amber-100/80 rounded-full"
                        ></div>
                    </div>
                </div>
                <div>
                    <p class="text-sm font-bold text-slate-900">Hangat Krem</p>
                    <p class="text-[10px] text-slate-500 font-medium mt-0.5">
                        Klasik & ramah
                    </p>
                </div>
            </button>
        </div>
    </div>
</div>

<!-- Modal / Popup Full Preview dengan Animasi Smooth Pop-up -->
{#if showPreview}
    <div
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6 font-sans"
    >
        <!-- Overlay dengan animasi fade-in -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div
            class="absolute inset-0 bg-slate-900/60 backdrop-blur-sm transition-opacity duration-300 opacity-100"
            onclick={() => (showPreview = false)}
        ></div>

        <!-- Konten Modal dengan animasi scale-up & smooth pop-up -->
        <div
            class="relative w-full max-w-3xl bg-slate-50 rounded-3xl shadow-2xl overflow-hidden transform transition-all duration-300 scale-100 opacity-100 flex flex-col max-h-[95vh]"
        >
            <!-- Header Modal -->
            <div
                class="px-8 py-5 border-b border-slate-200 bg-white flex items-center justify-between shrink-0"
            >
                <div>
                    <h3
                        class="text-lg font-black text-slate-900 tracking-tight"
                    >
                        Preview Tema: {previewTheme === "minimal"
                            ? "Minimal Putih"
                            : previewTheme === "gelap"
                              ? "Gelap Elegan"
                              : "Hangat Krem"}
                    </h3>
                    <p class="text-[11px] text-slate-400 font-medium mt-0.5">
                        Tampilan invoice yang akan dikirim ke klien Anda.
                    </p>
                </div>
                <button
                    onclick={() => (showPreview = false)}
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

            <!-- Area Tampilan Preview Faktur -->
            <div
                class="p-8 overflow-y-auto bg-slate-100 flex items-center justify-center"
            >
                <div
                    class="w-full max-w-xl rounded-2xl shadow-xl p-10 transition-colors duration-300 border
          {previewTheme === 'gelap'
                        ? 'bg-slate-900 text-white border-slate-700'
                        : previewTheme === 'krem'
                          ? 'bg-[#FDFBF7] text-slate-900 border-amber-200/60'
                          : 'bg-white text-slate-900 border-slate-200'}"
                >
                    <!-- Header Faktur -->
                    <div class="flex justify-between items-start mb-10">
                        <div>
                            <h2 class="text-2xl font-black tracking-tight">
                                Kilat Studio
                            </h2>
                            <p
                                class="text-xs mt-1 {previewTheme === 'gelap'
                                    ? 'text-slate-400'
                                    : 'text-slate-500'}"
                            >
                                hello@kilatstudio.id
                            </p>
                            <p
                                class="text-xs mt-0.5 {previewTheme === 'gelap'
                                    ? 'text-slate-400'
                                    : 'text-slate-500'}"
                            >
                                Jl. Sudirman No. 45, Jakarta
                            </p>
                        </div>
                        <div class="text-right">
                            <h1
                                class="text-3xl font-black tracking-widest uppercase {previewTheme ===
                                'gelap'
                                    ? 'text-white'
                                    : 'text-slate-200'}"
                            >
                                Invoice
                            </h1>
                            <p class="text-sm font-bold mt-2">INV-2026-092</p>
                            <p
                                class="text-[10px] font-medium uppercase tracking-wider mt-1 {previewTheme ===
                                'gelap'
                                    ? 'text-slate-400'
                                    : 'text-slate-500'}"
                            >
                                10 Oktober 2026
                            </p>
                        </div>
                    </div>

                    <!-- Ditagihkan ke -->
                    <div class="mb-10">
                        <p
                            class="text-[10px] font-bold uppercase tracking-wider mb-2 {previewTheme ===
                            'gelap'
                                ? 'text-slate-400'
                                : 'text-slate-500'}"
                        >
                            Ditagihkan Kepada
                        </p>
                        <p class="text-sm font-bold">PT Sinar Abadi</p>
                        <p
                            class="text-xs font-medium mt-1 {previewTheme ===
                            'gelap'
                                ? 'text-slate-400'
                                : 'text-slate-500'}"
                        >
                            Bapak Budi Santoso
                        </p>
                    </div>

                    <!-- Tabel Tagihan -->
                    <div class="w-full">
                        <div
                            class="flex justify-between border-t border-b py-3 mb-4 text-[10px] font-bold uppercase tracking-wider
              {previewTheme === 'gelap'
                                ? 'border-slate-700 text-slate-400'
                                : previewTheme === 'krem'
                                  ? 'border-amber-200/60 text-slate-500'
                                  : 'border-slate-200 text-slate-500'}"
                        >
                            <span>Deskripsi Layanan</span>
                            <span>Jumlah</span>
                        </div>

                        <div
                            class="flex justify-between text-sm font-bold mb-4"
                        >
                            <span>Jasa Pembuatan Website Company Profile</span>
                            <span>Rp 5.000.000</span>
                        </div>
                        <div
                            class="flex justify-between text-sm font-bold mb-8"
                        >
                            <span>Maintenance Server (1 Tahun)</span>
                            <span>Rp 1.500.000</span>
                        </div>

                        <!-- Total Akhir -->
                        <div
                            class="flex justify-between items-center border-t pt-5 mt-4
              {previewTheme === 'gelap'
                                ? 'border-slate-700'
                                : previewTheme === 'krem'
                                  ? 'border-amber-200/60'
                                  : 'border-slate-200'}"
                        >
                            <span
                                class="text-xs font-bold uppercase tracking-wider {previewTheme ===
                                'gelap'
                                    ? 'text-slate-400'
                                    : 'text-slate-500'}">Total Tagihan</span
                            >
                            <span class="text-2xl font-black">Rp 6.500.000</span
                            >
                        </div>
                    </div>
                </div>
            </div>

            <!-- Footer Action Modal -->
            <div
                class="px-8 py-5 border-t border-slate-200 bg-white flex justify-end gap-3 shrink-0"
            >
                <button
                    onclick={() => (showPreview = false)}
                    class="px-6 py-2.5 rounded-xl font-bold text-sm text-slate-600 bg-white border border-slate-200 hover:bg-slate-50 transition-colors"
                >
                    Tutup
                </button>
                <button
                    onclick={applyTheme}
                    class="px-6 py-2.5 rounded-xl font-bold text-sm text-white bg-slate-900 hover:bg-slate-800 shadow-lg shadow-slate-900/20 transition-all hover:-translate-y-0.5"
                >
                    Gunakan Tema Ini
                </button>
            </div>
        </div>
    </div>
{/if}
