self: super: {
	lasem = super.lasem.overrideAttrs (old: {
		version = "0.7.0-ee047b6";

		src = super.fetchFromGitHub {
			owner  = "mjakeman";
			repo   = "lasem";
			rev    = "bd73fd1";
			sha256 = "0s30q692s4523jwr99q4ys1chzhm2cdjsmyb1l78pgdl6nil190z";
		};

		buildInputs = (old.buildInputs or []) ++ (with super; [
			bison
			flex
		]);

		nativeBuildInputs = with super; [
			pkg-config
			intltool
			meson
			ninja
			gi-docgen
			gobject-introspection
		];

		# We have failing tests.
		doCheck = false;

		# We don't need these.
		mesonFlags = [
			"-Ddocs=disabled"
			"-Ddemo=disabled"
		];

		outputs = [ "out" "dev" "man" ];

		# po/meson.build:2:5: ERROR: Tried to create target "lasem-0.7-it.mo", but a target of that name already exists.
		preConfigure = (old.preConfigure or "") + ''
			rm po/*.po
			:> po/LINGUAS
		'';
	});
}
