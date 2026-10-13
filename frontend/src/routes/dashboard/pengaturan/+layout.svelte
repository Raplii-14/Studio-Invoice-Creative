<script>
    import { page } from "$app/stores";
    let { children } = $props();

    const tabs = [
        {
            name: "Akun",
            href: "/dashboard/pengaturan",
            icon: "M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z",
        },
        {
            name: "Profil Bisnis",
            href: "/dashboard/pengaturan/profil-bisnis",
            icon: "M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4",
        },
        {
            name: "Tampilan",
            href: "/dashboard/pengaturan/tampilan",
            icon: "M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01",
        },
        {
            name: "Pembayaran",
            href: "/dashboard/pengaturan/pembayaran",
            icon: "M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z",
        },
        {
            name: "Notifikasi",
            href: "/dashboard/pengaturan/notifikasi",
            icon: "M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9",
        },
    ];
</script>

<div class="p-10 max-w-7xl mx-auto w-full font-sans">
    <!-- Header Judul -->
    <div class="mb-8">
        <h1 class="text-2xl font-black text-slate-900 tracking-tight">
            Pengaturan
        </h1>
        <p class="text-sm text-slate-500 font-medium mt-0.5">
            Suaikan preferensi aplikasi
        </p>
    </div>

    <div class="flex flex-col lg:flex-row gap-8 items-start">
        <!-- Menu Tab Navigasi Samping (Kiri) -->
        <div class="w-full lg:w-64 bg-transparent shrink-0 space-y-1">
            {#each tabs as tab}
                {@const isActive = $page.url.pathname === tab.href}
                <a
                    href={tab.href}
                    class="flex items-center px-4 py-3.5 rounded-2xl font-bold text-xs transition-colors {isActive
                        ? 'bg-slate-900 text-white shadow-md shadow-slate-900/10'
                        : 'text-slate-600 hover:bg-slate-100/80'}"
                >
                    <svg
                        class="w-4 h-4 mr-3 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d={tab.icon}
                        />
                    </svg>
                    {tab.name}
                </a>
            {/each}
        </div>

        <!-- Kotak Konten Form (Kanan) dengan padding dalam p-10 yang lega -->
        <div
            class="flex-1 w-full bg-white rounded-3xl border border-slate-100 shadow-[0_8px_30px_rgb(0,0,0,0.04)] p-10"
        >
            {@render children()}
        </div>
    </div>
</div>
