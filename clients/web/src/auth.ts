import Keycloak, {KeycloakConfig, KeycloakInitOptions} from 'keycloak-js'

const host = "http://localhost:8080"
const clientID = "client"
const realm = "realm"

const initConfigs: KeycloakConfig = {
    url: host,
    realm: realm,
    clientId: clientID,
}

const initOptions: KeycloakInitOptions = {
    onLoad: "login-required",
    enableLogging: true
}

const keycloak = new Keycloak(initConfigs);

export default function StartApp(renderer: () => void) {
    keycloak.init(initOptions).then(authorized => {
            if (authorized) {
                setInterval(() => {
                    keycloak.updateToken(60).then((refreshed: any) => {
                        if (refreshed) {
                            console.log('Token refreshed' + refreshed);
                        }
                    }).catch(() => {
                        console.log('Failed to refresh token');
                    });
                }, 30000)
            }
        }
    ).catch((err) => {
        console.log(err)
    }).finally(() => {
        renderer()
    })
}

export function Logout() {
    keycloak.logout()
}

export function HasRole(role: string): boolean {
    return keycloak.hasResourceRole(role, clientID)
}

export function PreferredName(): string {
    return keycloak.idTokenParsed?.preferred_username ?? ""
}

export function Email(): string {
    return keycloak.idTokenParsed?.email ?? ""
}

export function AvatarURL(): string {
    return keycloak.idTokenParsed?.avatar_url ?? ""
}

export function GetIDToken(): string {
    return keycloak.idToken ?? ""
}

export function GetAccessToken(): string {
    return keycloak.token ?? ""
}
