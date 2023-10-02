// Copyright Nicolas Paul (2023)
//
// * Nicolas Paul
//
// This software is a computer program whose purpose is to allow the hosting
// and sharing of Go modules using a personal domain.
//
// This software is governed by the CeCILL license under French law and
// abiding by the rules of distribution of free software.  You can  use,
// modify and/ or redistribute the software under the terms of the CeCILL
// license as circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and  rights to copy,
// modify and redistribute granted by the license, users are provided only
// with a limited warranty  and the software's author,  the holder of the
// economic rights,  and the successive licensors  have only  limited
// liability.
//
// In this respect, the user's attention is drawn to the risks associated
// with loading,  using,  modifying and/or developing or reproducing the
// software by the user in light of its specific status of free software,
// that may mean  that it is complicated to manipulate,  and  that  also
// therefore means  that it is reserved for developers  and  experienced
// professionals having in-depth computer knowledge. Users are therefore
// encouraged to load and test the software's suitability as regards their
// requirements in conditions enabling the security of their systems and/or
// data to be ensured and,  more generally, to use and operate it in the
// same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

import { useRouter } from "next/router";
import { useConfig } from "nextra-theme-docs"

export default {
	docsRepositoryBase: "https://github.com/nc0fr/svgu/tree/main/doc",
	useNextSeoProps() {
		return {
			titleTemplate: "%s — SVGU"
		}
	},
	head: () => {
		const { frontMatter } = useConfig();
		const { asPath, defaultLocale, locale } = useRouter();

		// Default values.
		const url: string = "https://svgu.nc0.fr"
			+ (defaultLocale === locale ? asPath : `/${locale}${asPath}`);
		const title = "SVGU";
		const description = "SVGU is a tool to publish your Go modules on \
			your own domain.";
		const keywords: string[] = [ "svgu", "go", "golang", "module",
			"package", "github", "google" ];

		return (
			<>
				{ /* Basic metadata */ }
				<meta name="viewport" 
					content="width=device-width, initial-scale=1.0" />
				<meta name="author" content="Nicolas Paul" />
				<meta name="description" 
					content={frontMatter.description || description } />
				<meta name="keywords" content={keywords.join(",")} />
				<meta name="color-scheme" content="light dark" />
				<meta name="robots" content="all" />

				{ /* Open Graph Protocol */ }
				<meta name="og:title" property={frontMatter.title || title} />
				<meta name="og:type" property="website" />
				<meta name="og:url" property={url} />
				{ /* TODO: og:image:* */ }
				<meta name="og:description"
					property={frontMatter.description || description } />
				<meta name="og:determiner" property="" />
				<meta name="og:locale" property={locale} />
				<meta name="og:site_name" property="nc0.fr" />

				{ /* Twitter Cards */ }
				<meta name="twitter:card" content="summary" />
				<meta name="twitter:site" content="nc0fr" />
				<meta name="twitter:title" 
					content={frontMatter.title || title} />
				<meta name="twitter:description"
					content={frontMatter.description || description } />
				{ /* TODO: twitter:image and twitter:image:alt*/ }
			</>
		);
	},
	darkMode: true,
	nextThemes: {
		defaultTheme: "system",
		enableTheme: true,
		enableColorScheme: true,
		themes: [ "light", "dark" ]
	},
	logo: <span>SVGU</span>,
	logoLink: true,
	project: {
		link: "https://github.com/nc0fr/svgu",
		// No need for an icon as the default is the GitHub one.
		// See: 
		// https://nextra.site/docs/docs-theme/theme-configuration#project-link
	},
	chat: {
		link: 'https://twitter.com/nc0_fr',
		icon: (
			<svg width="24" height="24" viewBox="0 0 248 204">
				<path
					fill="currentColor"
					d="M221.95 51.29c.15 2.17.15 4.34.15 6.53 0 66.73-50.8 143.69-143.69 143.69v-.04c-27.44.04-54.31-7.82-77.41-22.64 3.99.48 8 .72 12.02.73 22.74.02 44.83-7.61 62.72-21.66-21.61-.41-40.56-14.5-47.18-35.07a50.338 50.338 0 0 0 22.8-.87C27.8 117.2 10.85 96.5 10.85 72.46v-.64a50.18 50.18 0 0 0 22.92 6.32C11.58 63.31 4.74 33.79 18.14 10.71a143.333 143.333 0 0 0 104.08 52.76 50.532 50.532 0 0 1 14.61-48.25c20.34-19.12 52.33-18.14 71.45 2.19 11.31-2.23 22.15-6.38 32.07-12.26a50.69 50.69 0 0 1-22.2 27.93c10.01-1.18 19.79-3.86 29-7.95a102.594 102.594 0 0 1-25.2 26.16z"
				/>
			</svg>
		)
	},
	banner: {
		dismissible: true,
		text: (<span>👋 Welcome to the new SVGU documentation!</span>)
	},
	direction: "ltr",
	navigation: true,
	footer: {
		text: (
			<span>
				Copyright 2023 <a href="https://nc0.fr" target="_blank">
				Nicolas Paul</a> 
				<br /> Creative Commons Attribution-ShareAlike 4.0.
			</span>
		)
	}
}
